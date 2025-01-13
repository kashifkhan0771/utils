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
}
