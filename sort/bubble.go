package sort

// BubbleSort performs an in-place bubble sort on a slice of numeric values.
// It implements the bubble sort algorithm which repeatedly steps through the slice,
// compares adjacent elements and swaps them if they are in the wrong order.
// Time Complexity: O(n²) where n is the length of the slice
// Space Complexity: O(1) as it sorts in-place
func BubbleSort[T number](arr []T) []T {
	n := len(arr)
	for i := range n - 1 {
		for j := range n - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
