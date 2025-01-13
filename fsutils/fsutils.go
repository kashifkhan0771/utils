package fsutils

import "fmt"

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
