package sort

// SelectionSort performs an in-place selection sort on a slice of numeric values.
// It works by dividing the input array into a sorted and an unsorted region.
// In each iteration, it finds the minimum element from the unsorted region
// and places it at the beginning of the sorted region.
// Time Complexity:
//   - Best Case: O(n²)
//   - Average Case: O(n²)
//   - Worst Case: O(n²)
//
// Space Complexity: O(1) as it sorts in-place
func SelectionSort[T number](arr []T) []T {
	n := len(arr)
	for i := range n - 1 {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}
