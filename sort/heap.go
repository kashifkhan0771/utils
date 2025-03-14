package sort

// HeapSort performs heap sort on a slice of numbers.
// It takes a slice of any numeric type T and returns a sorted slice.
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func HeapSort[T number](arr []T) []T {
	return heapSort(arr)
}

// heapSort is an internal implementation of heap sort algorithm.
// It first builds a max heap and then repeatedly extracts the maximum element.
func heapSort[T number](arr []T) []T {
	n := len(arr)
	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// Extract elements from heap one by one
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // Move current root to end
		heapify(arr, i, 0)              // Call max heapify on the reduced heap
	}
	return arr
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
