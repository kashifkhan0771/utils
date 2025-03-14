package sort

// InsertionSort performs an in-place insertion sort on a slice of numeric values.
// It works by building a sorted array one element at a time by repeatedly taking
// the next unsorted element and inserting it into its correct position in the sorted portion.
// Time Complexity:
//   - Best Case: O(n) when array is already sorted
//   - Average Case: O(n²)
//   - Worst Case: O(n²) when array is reverse sorted
//
// Space Complexity: O(1) as it sorts in-place
func InsertionSort[T number](arr []T) []T {
	n := len(arr)
	for i := range n - 1 {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}
