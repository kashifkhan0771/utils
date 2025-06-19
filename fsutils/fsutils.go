package fsutils

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
	"time"
)

type ByteSize int64

const (
	KB ByteSize = 1024
	MB ByteSize = KB * 1024
	GB ByteSize = MB * 1024
	TB ByteSize = GB * 1024
)

// FormatFileSize formats a file size given in bytes into a human-readable string
// with appropriate units (B, KB, MB, GB, TB).
func FormatFileSize(size int64) string {
	switch {
	case size >= int64(TB):
		return fmt.Sprintf("%.2f TB", float64(size)/float64(TB))
	case size >= int64(GB):
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= int64(MB):
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= int64(KB):
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

// FindFiles searches for files with the specified extension in the given root directory
// and returns a slice of matching file paths.
func FindFiles(root string, extension string) ([]string, error) {
	root = filepath.Clean(root)
	files := make([]string, 0)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}

		if !info.IsDir() && (filepath.Ext(path) == extension || extension == "") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// FindFilesWithFilter searches for files in the given root directory using a custom filter function.
// The filterFn parameter allows you to define arbitrary logic for file selection (e.g., size, mod time, type).
// If filterFn is nil, all non-directory, non-symlink files are returned.
func FindFilesWithFilter(root string, filterFn func(fs.FileInfo) bool) ([]string, error) {
	root = filepath.Clean(root)
	files := make([]string, 0)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip symlinks for safety and consistency
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}

		// If filterFn is nil, default to non-directory, non-symlink files
		if filterFn == nil {
			if !info.IsDir() {
				files = append(files, path)
			}

			return nil
		}

		if filterFn(info) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// GetDirectorySize calculates the total size (in bytes) of all files within the specified directory
func GetDirectorySize(path string) (int64, error) {
	var size int64 = 0

	err := filepath.Walk(path, func(fPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return size, nil
}

// FilesIdentical compares two files byte by byte to determine if they are identical
func FilesIdentical(file1, file2 string) (bool, error) {
	f1, err := os.Open(file1)
	if err != nil {
		return false, err
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	const chunkSize = 32 * 1024
	b1 := make([]byte, chunkSize)
	b2 := make([]byte, chunkSize)

	for {
		n1, err1 := f1.Read(b1)
		n2, err2 := f2.Read(b2)

		if n1 != n2 || !bytes.Equal(b1[:n1], b2[:n2]) {
			return false, nil
		}

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			}

			if err1 == io.EOF || err2 == io.EOF {
				return false, nil
			}

			return false, fmt.Errorf("error reading files: %v, %v", err1, err2)
		}
	}
}

// DirsIdentical compares two directories to determine if they are identical
// It returns true if both directories contain the same files with identical content,
// and false otherwise
func DirsIdentical(dir1, dir2 string) (bool, error) {
	dir1 = filepath.Clean(dir1)
	dir2 = filepath.Clean(dir2)

	// Check if either path is a symlink
	for _, dir := range []string{dir1, dir2} {
		info, err := os.Lstat(dir)
		if err != nil {
			return false, err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return false, fmt.Errorf("symlinks are not supported: %s", dir)
		}
	}

	files1, err := FindFiles(dir1, "")
	if err != nil {
		return false, err
	}

	files2, err := FindFiles(dir2, "")
	if err != nil {
		return false, err
	}

	if len(files1) != len(files2) {
		return false, nil
	}

	type result struct {
		path      string
		identical bool
		err       error
	}

	workers := make(chan struct{}, 10)
	results := make(chan result)

	for _, file1 := range files1 {
		relativePath1, err := filepath.Rel(dir1, file1)
		if err != nil {
			return false, err
		}

		file2 := filepath.Join(dir2, relativePath1)
		if _, err := os.Lstat(file2); os.IsNotExist(err) {
			return false, nil
		}

		workers <- struct{}{}
		go func(f1 string, f2 string, rPath string) {
			defer func() { <-workers }()
			identical, err := FilesIdentical(f1, f2)
			results <- result{f1, identical, err}
		}(file1, file2, relativePath1)
	}

	matched := make(map[string]bool)
	for range files1 {
		r := <-results
		if r.err != nil {
			return false, r.err
		}

		if !r.identical {
			return false, nil
		}

		matched[r.path] = true
	}

	return len(matched) == len(files1), nil
}

type FileMetadata struct {
	Name    string      `json:"name"`
	Size    int64       `json:"size"`
	IsDir   bool        `json:"is_dir"`
	ModTime time.Time   `json:"mod_time"`
	Mode    os.FileMode `json:"mode"`
	Path    string      `json:"path"`
	Ext     string      `json:"ext"`
	Owner   string      `json:"owner"`
}

// GetFileMetadata retrieves metadata for the specified file path.
// It returns a FileMetadata struct containing details about the file.
func GetFileMetadata(filePath string) (FileMetadata, error) {
	if filePath == "" {
		return FileMetadata{}, fmt.Errorf("file path cannot be empty")
	}

	filePath = filepath.Clean(filePath)

	info, err := os.Lstat(filePath)
	if err != nil {
		return FileMetadata{}, err
	}

	path, err := filepath.Abs(filePath)
	if err != nil {
		return FileMetadata{}, err
	}

	metadata := FileMetadata{
		Name:    info.Name(),
		Size:    info.Size(),
		IsDir:   info.IsDir(),
		ModTime: info.ModTime(),
		Mode:    info.Mode(),
		Path:    path,
		Ext:     filepath.Ext(info.Name()),
	}

	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		owner, err := user.LookupId(fmt.Sprint(stat.Uid))
		if err != nil {
			return metadata, fmt.Errorf("failed to lookup owner: %w", err)
		}

		metadata.Owner = owner.Username
	} else {
		metadata.Owner = "unknown"
	}

	return metadata, nil
}
