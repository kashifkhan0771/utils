package slice

// RemoveDuplicateStr removes duplicate strings and returns a new slice without duplicates, preserving input order.
func RemoveDuplicateStr(strSlice []string) []string {
	return unique(strSlice)
}

// RemoveDuplicateInt removes duplicate integers and returns a new slice without duplicates, preserving input order.
func RemoveDuplicateInt(intSlice []int) []int {
	return unique(intSlice)
}

func unique[T comparable](in []T) []T {
	if in == nil {
		return nil
	}

	seen := make(map[T]struct{}, len(in))
	out := make([]T, 0, len(in))

	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}

	return out
}
