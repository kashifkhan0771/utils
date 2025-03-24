package sort

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"
)

func generateRandomSliceInt(size int) []int {
	max := big.NewInt(1000)
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		slice[i] = int(num.Int64())
	}

	return slice
}

func generateRandomSliceFloat(size int) []float64 {
	maxUint64 := new(big.Int).Lsh(big.NewInt(1), 53) // 2^53 for float64 precision
	slice := make([]float64, size)

	for i := 0; i < size; i++ {
		n, err := rand.Int(rand.Reader, maxUint64)
		if err != nil {
			panic(err)
		}
		// Convert to float64 in range [0, 1), then scale to [0, 1000)
		f := float64(n.Int64()) / float64(1<<53)
		slice[i] = f * 1000
	}

	return slice
}

func isSorted[T number](arr []T) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

type Fn[T number] struct {
	Name string
	Fn   func([]T) []T
}

func sortFns[T number]() []Fn[T] {
	return []Fn[T]{
		{
			Name: "BubbleSort",
			Fn:   BubbleSort[T],
		},
		{
			Name: "SelectionSort",
			Fn:   SelectionSort[T],
		},
		{
			Name: "InsertionSort",
			Fn:   InsertionSort[T],
		},
		{
			Name: "MergeSort",
			Fn:   MergeSort[T],
		},
		{
			Name: "QuickSort",
			Fn:   QuickSort[T],
		},
		{
			Name: "HeapSort",
			Fn:   HeapSort[T],
		},
	}
}

func TestSortInt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arr  []int
	}{
		{
			name: "success - sorted array",
			arr:  []int{1, 2, 3, 4, 5},
		},
		{
			name: "success - array with duplicates",
			arr:  []int{1, 2, 2, 3, 4},
		},
		{
			name: "success - unsorted array",
			arr:  generateRandomSliceInt(5),
		},
		{
			name: "success - empty array",
			arr:  []int{},
		},
		{
			name: "success - nil array",
			arr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for _, fn := range sortFns[int]() {
				t.Run(fn.Name, func(t *testing.T) {
					t.Parallel()
					inputCopy := make([]int, len(tt.arr))
					copy(inputCopy, tt.arr)
					fn.Fn(inputCopy)
					if !isSorted(inputCopy) {
						t.Errorf("%v is not sorted", inputCopy)
					}
				})
			}
		})
	}
}

func TestSortFloat(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arr  []float64
	}{
		{
			name: "success - sorted array",
			arr:  []float64{1.1, 2.2, 3.3, 4.4, 5.5},
		},
		{
			name: "success - unsorted array",
			arr:  generateRandomSliceFloat(5),
		},
		{
			name: "success - empty array",
			arr:  []float64{},
		},
		{
			name: "success - nil array",
			arr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for _, fn := range sortFns[float64]() {
				t.Run(fn.Name, func(t *testing.T) {
					t.Parallel()
					inputCopy := make([]float64, len(tt.arr))
					copy(inputCopy, tt.arr)
					fn.Fn(inputCopy)
					if !isSorted(inputCopy) {
						t.Errorf("%v is not sorted", inputCopy)
					}
				})
			}
		})
	}
}

func BenchmarkSort(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"Small", 10},
		{"Medium", 100},
		{"Large", 1000},
	}

	for _, bm := range benchmarks {
		input := generateRandomSliceInt(bm.size)
		for _, fn := range sortFns[int]() {
			b.Run(fn.Name+"-"+bm.name+"("+strconv.Itoa(bm.size)+")", func(b *testing.B) {
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					// Create a fresh copy of the input for each iteration
					inputCopy := make([]int, len(input))
					copy(inputCopy, input)
					fn.Fn(inputCopy)
				}
			})
		}
	}
}
