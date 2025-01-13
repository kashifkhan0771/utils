package fsutils

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
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
	files := make([]string, 0)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
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

	const chunkSize = 4096
	b1 := make([]byte, chunkSize)
	b2 := make([]byte, chunkSize)

	for {
		_, err1 := f1.Read(b1)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			}
			
			return false, fmt.Errorf("error reading files: %v, %v", err1, err2)
		}

		if !bytes.Equal(b1, b2) {
			return false, nil
		}
	}
}
