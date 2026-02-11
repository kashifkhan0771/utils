package pdf

import (
	"fmt"
	"strconv"
	"strings"
)

// parsePageRange parses a page range string like "1-3" or "5" and returns
// the start and end page numbers (1-indexed). For a single page like "5",
// both start and end will be 5.
func parsePageRange(rangeStr string) (int, int, error) {
	rangeStr = strings.TrimSpace(rangeStr)
	if rangeStr == "" {
		return 0, 0, fmt.Errorf("empty page range")
	}

	parts := strings.SplitN(rangeStr, "-", 2)

	start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start page %q: %w", parts[0], err)
	}

	if start < 1 {
		return 0, 0, fmt.Errorf("start page must be >= 1, got %d", start)
	}

	end := start
	if len(parts) == 2 {
		end, err = strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return 0, 0, fmt.Errorf("invalid end page %q: %w", parts[1], err)
		}
	}

	if end < 1 {
		return 0, 0, fmt.Errorf("end page must be >= 1, got %d", end)
	}

	if end < start {
		return 0, 0, fmt.Errorf("end page (%d) must be >= start page (%d)", end, start)
	}

	return start, end, nil
}
