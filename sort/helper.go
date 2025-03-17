package sort

// mergeSort recursively divides the array into two halves and merges them in sorted order.
// Parameters:
//   - arr: The input array being sorted
//   - left: Starting index of the subarray
//   - right: Ending index of the subarray
func mergeSort[T number](arr []T, left, right int) []T {
	if left < right {
		mid := left + (right-left)/2
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
	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}
	for i := 0; i < n2; i++ {
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

// quickSort recursively sorts a portion of the array between left and right indices.
// Parameters:
//   - arr: The input array being sorted
//   - left: Starting index of the subarray
//   - right: Ending index of the subarray
func quickSort[T number](arr []T, left, right int) []T {
	if left < right {
		pivot := partition(arr, left, right)
		_ = quickSort(arr, left, pivot-1)
		_ = quickSort(arr, pivot+1, right)
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

// heapify maintains the heap property for the subtree rooted at index i.
// It ensures that the largest element is at the root of the subtree.
// Parameters:
//   - arr: the input array being heapified
//   - n: size of the heap
//   - i: root index of the subtree
func heapify[T number](arr []T, n, i int) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // Left child
	right := 2*i + 2 // Right child

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // Swap
		heapify(arr, n, largest)                    // Recursively heapify the affected sub-tree
	}
}
