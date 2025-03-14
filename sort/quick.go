package sort

// QuickSort performs a quick sort on a slice of numeric values.
// It uses the divide-and-conquer strategy by selecting a 'pivot' element
// and partitioning the array around it such that smaller elements are on the left
// and larger elements are on the right.
// Time Complexity:
//   - Average Case: O(n log n)
//   - Worst Case: O(n²) when array is already sorted
//   - Best Case: O(n log n)
//
// Space Complexity: O(log n) due to recursive call stack
func QuickSort[T number](arr []T) []T {
	return quickSort(arr, 0, len(arr)-1)
}

// quickSort recursively sorts a portion of the array between left and right indices.
// Parameters:
//   - arr: The input array being sorted
//   - left: Starting index of the subarray
//   - right: Ending index of the subarray
func quickSort[T number](arr []T, left, right int) []T {
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
	return arr
}

// partition rearranges the array segment and returns the pivot position.
// It selects the rightmost element as pivot and places:
//   - all elements smaller than pivot to its left
//   - all elements larger than pivot to its right
//
// Parameters:
//   - arr: The input array being partitioned
//   - left: Starting index of the segment
//   - right: Ending index of the segment
//
// Returns:
//   - The final position of the pivot element
func partition[T number](arr []T, left, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
