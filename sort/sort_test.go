package sort

import "math/rand"

func generateRandomSliceInt(size int) []int {
	slice := make([]int, size)
	for i := range size {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func generateRandomSliceFloat(size int) []float64 {
	slice := make([]float64, size)
	for i := range size {
		slice[i] = rand.Float64() * 1000
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
