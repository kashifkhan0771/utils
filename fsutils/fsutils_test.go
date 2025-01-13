package fsutils

import (
	"fmt"
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
