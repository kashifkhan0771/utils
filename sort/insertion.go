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
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}
