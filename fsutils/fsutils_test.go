package fsutils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestFormatFileSize(t *testing.T) {
	tests := []struct {
		size     int64
		expected string
	}{
		{size: 0, expected: "0 B"},
		{size: 1, expected: "1 B"},
		{size: 512, expected: "512 B"},
		{size: 1023, expected: "1023 B"},
		{size: 1024, expected: "1.00 KB"},
		{size: 1536, expected: "1.50 KB"},
		{size: 2048, expected: "2.00 KB"},
		{size: 1048576, expected: "1.00 MB"},
		{size: 1572864, expected: "1.50 MB"},
		{size: 2097152, expected: "2.00 MB"},
		{size: 536870912, expected: "512.00 MB"},
		{size: 268435456, expected: "256.00 MB"},
		{size: 134217728, expected: "128.00 MB"},
		{size: 1073741824, expected: "1.00 GB"},
		{size: 1610612736, expected: "1.50 GB"},
		{size: 2147483648, expected: "2.00 GB"},
		{size: 1099511627776, expected: "1.00 TB"},
		{size: 1649267441664, expected: "1.50 TB"},
		{size: 2199023255552, expected: "2.00 TB"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d bytes", tt.size), func(t *testing.T) {
			result := FormatFileSize(tt.size)
			if result != tt.expected {
				t.Errorf("FormatFileSize(%d) = %s; want %s", tt.size, result, tt.expected)
			}
		})
	}
}

func TestFindFiles(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tempDir)

	// Create some test files
	files := []struct {
		path     string
		contents string
	}{
		{path: filepath.Join(tempDir, "file1.txt"), contents: "file1"},
		{path: filepath.Join(tempDir, "file2.txt"), contents: "file2"},
		{path: filepath.Join(tempDir, "file3.log"), contents: "file3"},
		{path: filepath.Join(tempDir, "file4.txt"), contents: "file4"},
		{path: filepath.Join(tempDir, "file5.md"), contents: "file5"},
		{path: filepath.Join(tempDir, "file6.md"), contents: "file6"},
		{path: filepath.Join(tempDir, "file7.log"), contents: "file7"},
		{path: filepath.Join(tempDir, "file8.txt"), contents: "file8"},
	}

	for _, file := range files {
		if err := os.WriteFile(file.path, []byte(file.contents), 0644); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		extension string
		expected  []string
	}{
		{extension: ".txt", expected: []string{
			filepath.Join(tempDir, "file1.txt"),
			filepath.Join(tempDir, "file2.txt"),
			filepath.Join(tempDir, "file4.txt"),
			filepath.Join(tempDir, "file8.txt"),
		}},
		{extension: ".log", expected: []string{
			filepath.Join(tempDir, "file3.log"),
			filepath.Join(tempDir, "file7.log"),
		}},
		{extension: ".md", expected: []string{
			filepath.Join(tempDir, "file5.md"),
			filepath.Join(tempDir, "file6.md"),
		}},
		{extension: ".json", expected: []string{}},
		{extension: "", expected: []string{
			filepath.Join(tempDir, "file1.txt"),
			filepath.Join(tempDir, "file2.txt"),
			filepath.Join(tempDir, "file3.log"),
			filepath.Join(tempDir, "file4.txt"),
			filepath.Join(tempDir, "file5.md"),
			filepath.Join(tempDir, "file6.md"),
			filepath.Join(tempDir, "file7.log"),
			filepath.Join(tempDir, "file8.txt"),
		}},
	}

	for _, tt := range tests {
		t.Run(tt.extension, func(t *testing.T) {
			result, err := FindFiles(tempDir, tt.extension)
			if err != nil {
				t.Fatalf("FindFiles() error = %v", err)
			}

			if len(result) != len(tt.expected) {
				t.Fatalf("FindFiles() = %v; want %v", result, tt.expected)
			}

			for i, file := range result {
				if file != tt.expected[i] {
					t.Errorf("FindFiles() = %v; want %v", result, tt.expected)
				}
			}
		})
	}

	t.Run("invalid path", func(t *testing.T) {
		_, err := FindFiles("/nonexistent/path", ".txt")
		if err == nil {
			t.Error("Expected error for nonexistent path")
		}
	})

	t.Run("permission denied", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "testdir")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tempDir)

		if err := os.Chmod(tempDir, 0000); err != nil {
			t.Fatal(err)
		}

		defer func() {
			if err := os.Chmod(tempDir, 0755); err != nil {
				t.Errorf("Failed to restore directory permissions: %v", err)
			}
		}()

		_, err = FindFiles(tempDir, ".txt")
		if err == nil {
			t.Error("Expected permission denied error")
		}
	})
}

func TestGetDirectorySize(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tempDir)

	// Create some test files with known sizes
	files := []struct {
		path     string
		contents string
		size     int64
	}{
		{path: filepath.Join(tempDir, "file1.txt"), contents: "file1", size: 5},
		{path: filepath.Join(tempDir, "file2.txt"), contents: "file2", size: 5},
		{path: filepath.Join(tempDir, "file3.log"), contents: "file3", size: 5},
		{path: filepath.Join(tempDir, "file4.txt"), contents: "file4", size: 5},
		{path: filepath.Join(tempDir, "file5.md"), contents: "file5", size: 5},
	}

	var expectedSize int64 = 0
	for _, file := range files {
		if err := os.WriteFile(file.path, []byte(file.contents), 0644); err != nil {
			t.Fatal(err)
		}
		expectedSize += file.size
	}

	t.Run("Calculate directory size", func(t *testing.T) {
		result, err := GetDirectorySize(tempDir)
		if err != nil {
			t.Fatalf("GetDirectorySize() error = %v", err)
		}

		if result != expectedSize {
			t.Errorf("GetDirectorySize() = %d; want %d", result, expectedSize)
		}
	})
}

func TestFilesIdentical(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tempDir)

	// Create some test files
	files := []struct {
		path     string
		contents string
	}{
		{path: filepath.Join(tempDir, "file1.txt"), contents: "file1"},
		{path: filepath.Join(tempDir, "file2.txt"), contents: "file1"},
		{path: filepath.Join(tempDir, "file3.txt"), contents: "file3"},
		{path: filepath.Join(tempDir, "file4.txt"), contents: "file4"},
	}

	for _, file := range files {
		if err := os.WriteFile(file.path, []byte(file.contents), 0644); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		file1    string
		file2    string
		expected bool
	}{
		{file1: filepath.Join(tempDir, "file1.txt"), file2: filepath.Join(tempDir, "file2.txt"), expected: true},
		{file1: filepath.Join(tempDir, "file1.txt"), file2: filepath.Join(tempDir, "file3.txt"), expected: false},
		{file1: filepath.Join(tempDir, "file3.txt"), file2: filepath.Join(tempDir, "file4.txt"), expected: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s vs %s", tt.file1, tt.file2), func(t *testing.T) {
			result, err := FilesIdentical(tt.file1, tt.file2)
			if err != nil {
				t.Fatalf("FilesIdentical() error = %v", err)
			}

			if result != tt.expected {
				t.Errorf("FilesIdentical() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestDirsIdentical(t *testing.T) {
	// Create temporary directories for testing
	tempDir1, err := os.MkdirTemp("", "testdir1")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir1)

	tempDir2, err := os.MkdirTemp("", "testdir2")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir2)

	// Create some test files in both directories
	files1 := []struct {
		path     string
		contents string
	}{
		{path: filepath.Join(tempDir1, "file1.txt"), contents: "file1"},
		{path: filepath.Join(tempDir1, "file2.txt"), contents: "file2"},
		{path: filepath.Join(tempDir1, "file3.log"), contents: "file3"},
	}

	files2 := []struct {
		path     string
		contents string
	}{
		{path: filepath.Join(tempDir2, "file1.txt"), contents: "file1"},
		{path: filepath.Join(tempDir2, "file2.txt"), contents: "file2"},
		{path: filepath.Join(tempDir2, "file3.log"), contents: "file3"},
	}

	for _, file := range files1 {
		if err := os.WriteFile(file.path, []byte(file.contents), 0644); err != nil {
			t.Fatal(err)
		}
	}

	for _, file := range files2 {
		if err := os.WriteFile(file.path, []byte(file.contents), 0644); err != nil {
			t.Fatal(err)
		}
	}

	t.Run("Identical directories", func(t *testing.T) {
		result, err := DirsIdentical(tempDir1, tempDir2)
		if err != nil {
			t.Fatalf("DirsIdentical() error = %v", err)
		}

		if !result {
			t.Errorf("DirsIdentical() = %v; want %v", result, true)
		}
	})

	if err := os.WriteFile(filepath.Join(tempDir2, "file2.txt"), []byte("modified"), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("Non-identical directories", func(t *testing.T) {
		result, err := DirsIdentical(tempDir1, tempDir2)
		if err != nil {
			t.Fatalf("DirsIdentical() error = %v", err)
		}

		if result {
			t.Errorf("DirsIdentical() = %v; want %v", result, false)
		}
	})

	if err := os.WriteFile(filepath.Join(tempDir2, "file4.txt"), []byte("file4"), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("Directories with different number of files", func(t *testing.T) {
		result, err := DirsIdentical(tempDir1, tempDir2)
		if err != nil {
			t.Fatalf("DirsIdentical() error = %v", err)
		}

		if result {
			t.Errorf("DirsIdentical() = %v; want %v", result, false)
		}
	})

	t.Run("nested directories", func(t *testing.T) {
		dir1, dir2 := setupNestedDirs(t)
		defer os.RemoveAll(dir1)
		defer os.RemoveAll(dir2)

		identical, err := DirsIdentical(dir1, dir2)
		if err != nil {
			t.Fatal(err)
		}
		if !identical {
			t.Error("Nested directories should be identical")
		}
	})
}

func setupNestedDirs(t *testing.T) (string, string) {
	// Create temporary directories for testing
	tempDir1, err := os.MkdirTemp("", "nestedtestdir1")
	if err != nil {
		t.Fatal(err)
	}

	tempDir2, err := os.MkdirTemp("", "nestedtestdir2")
	if err != nil {
		t.Fatal(err)
	}

	// Create nested directory structure in both directories
	nestedDirs := []string{
		"dir1",
		"dir1/dir2",
		"dir1/dir2/dir3",
	}

	for _, dir := range nestedDirs {
		if err := os.MkdirAll(filepath.Join(tempDir1, dir), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.MkdirAll(filepath.Join(tempDir2, dir), 0755); err != nil {
			t.Fatal(err)
		}
	}

	// Create some test files in the nested directories
	files := []struct {
		path     string
		contents string
	}{
		{path: filepath.Join(tempDir1, "dir1/file1.txt"), contents: "file1"},
		{path: filepath.Join(tempDir1, "dir1/dir2/file2.txt"), contents: "file2"},
		{path: filepath.Join(tempDir1, "dir1/dir2/dir3/file3.txt"), contents: "file3"},
	}

	for _, file := range files {
		if err := os.WriteFile(file.path, []byte(file.contents), 0644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(tempDir2, file.path[len(tempDir1):]), []byte(file.contents), 0644); err != nil {
			t.Fatal(err)
		}
	}

	return tempDir1, tempDir2
}

func TestGetFileMetadata(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test file
	filePath := filepath.Join(tempDir, "testfile.txt")
	contents := "test file contents"
	if err := os.WriteFile(filePath, []byte(contents), 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("Get file metadata", func(t *testing.T) {
		metadata, err := GetFileMetadata(filePath)
		if err != nil {
			t.Fatalf("GetFileMetadata() error = %v", err)
		}

		if metadata.Name != "testfile.txt" {
			t.Errorf("GetFileMetadata() Name = %v; want %v", metadata.Name, "testfile.txt")
		}

		if metadata.Size != int64(len(contents)) {
			t.Errorf("GetFileMetadata() Size = %v; want %v", metadata.Size, len(contents))
		}

		if metadata.IsDir {
			t.Errorf("GetFileMetadata() IsDir = %v; want %v", metadata.IsDir, false)
		}

		if metadata.Ext != ".txt" {
			t.Errorf("GetFileMetadata() Ext = %v; want %v", metadata.Ext, ".txt")
		}

		if metadata.Path != filePath {
			t.Errorf("GetFileMetadata() Path = %v; want %v", metadata.Path, filePath)
		}

		if metadata.Owner == "" {
			t.Errorf("GetFileMetadata() Owner should not be empty")
		}
	})

	t.Run("Get directory metadata", func(t *testing.T) {
		metadata, err := GetFileMetadata(tempDir)
		if err != nil {
			t.Fatalf("GetFileMetadata() error = %v", err)
		}

		if metadata.Name != filepath.Base(tempDir) {
			t.Errorf("GetFileMetadata() Name = %v; want %v", metadata.Name, filepath.Base(tempDir))
		}

		if !metadata.IsDir {
			t.Errorf("GetFileMetadata() IsDir = %v; want %v", metadata.IsDir, true)
		}
	})

	t.Run("Nonexistent file", func(t *testing.T) {
		_, err := GetFileMetadata(filepath.Join(tempDir, "nonexistent.txt"))
		if err == nil {
			t.Error("Expected error for nonexistent file")
		}
	})

	t.Run("Empty path", func(t *testing.T) {
		_, err := GetFileMetadata("")
		if err == nil {
			t.Error("Expected error for empty path")
		}
	})

	t.Run("Symlink", func(t *testing.T) {
		tempDir, err := os.MkdirTemp("", "testdir")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tempDir)

		// Create a file and a symlink to it
		filePath := filepath.Join(tempDir, "testfile.txt")
		if err := os.WriteFile(filePath, []byte("test"), 0644); err != nil {
			t.Fatal(err)
		}

		linkPath := filepath.Join(tempDir, "testlink")
		if err := os.Symlink(filePath, linkPath); err != nil {
			t.Fatal(err)
		}

		metadata, err := GetFileMetadata(linkPath)
		if err != nil {
			t.Fatal(err)
		}

		if metadata.Mode&os.ModeSymlink == 0 {
			t.Error("Expected symlink mode")
		}
	})
}
