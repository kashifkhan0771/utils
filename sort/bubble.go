package sort

// BubbleSort performs an in-place bubble sort on a slice of numeric values.
// It implements the bubble sort algorithm which repeatedly steps through the slice,
// compares adjacent elements and swaps them if they are in the wrong order.
// The function doesn't return new slice, but it sorts the slice in-place.
// Time Complexity: O(n²) where n is the length of the slice
// Space Complexity: O(1) as it sorts in-place
func BubbleSort[T number](arr []T) []T {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
