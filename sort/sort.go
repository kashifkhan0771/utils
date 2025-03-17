package sort

// number is a type constraint that matches all numeric types (integers and floats).
type number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

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
	for i := 0; i < n-1; i++ {
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

// MergeSort performs a merge sort on a slice of numeric values.
// It uses the divide-and-conquer strategy by recursively dividing the input array
// into two halves, sorting them, and then merging the sorted halves.
// Time Complexity: O(n log n) for all cases
// Space Complexity: O(n) for temporary arrays during merging
func MergeSort[T number](arr []T) []T {
	return mergeSort(arr, 0, len(arr)-1)
}

// QuickSort performs a quick sort on a slice of numeric values.
// It uses the divide-and-conquer strategy by selecting a 'pivot' element
// and partitioning the surrounding array such that smaller elements are on the left
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

// HeapSort performs heap sort on a slice of numbers.
// It takes a slice of any numeric type T and returns a sorted slice.
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func HeapSort[T number](arr []T) []T {
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
