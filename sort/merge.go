package sort

// MergeSort performs a merge sort on a slice of numeric values.
// It uses the divide-and-conquer strategy by recursively dividing the input array
// into two halves, sorting them, and then merging the sorted halves.
// Time Complexity: O(n log n) for all cases
// Space Complexity: O(n) for temporary arrays during merging
func MergeSort[T number](arr []T) []T {
	return mergeSort(arr, 0, len(arr)-1)
}

// mergeSort recursively divides the array into two halves and merges them in sorted order.
// Parameters:
//   - arr: The input array being sorted
//   - left: Starting index of the subarray
//   - right: Ending index of the subarray
func mergeSort[T number](arr []T, left, right int) []T {
	if left < right {
		mid := (left + right) / 2
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
	return arr
}

// merge combines two sorted subarrays into a single sorted array.
// Parameters:
//   - arr: The input array containing the subarrays to merge
//   - left: Starting index of the left subarray
//   - mid: Ending index of the left subarray
//   - right: Ending index of the right subarray
func merge[T number](arr []T, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid
	L := make([]T, n1)
	R := make([]T, n2)
	for i := range n1 {
		L[i] = arr[left+i]
	}
	for i := range n2 {
		R[i] = arr[mid+1+i]
	}
	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}
