package sort

import (
	"math"
	"testing"
)

func testBubbleSortInt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success - sorted array",
			args: args{[]int{1, 2, 3, 4, 5}},
		},
		{
			name: "success - unsorted array",
			args: args{generateRandomSliceInt(5)},
		},
		{
			name: "success - empty array",
			args: args{[]int{}},
		},
		{
			name: "success - nil array",
			args: args{nil},
		},
		{
			name: "int edge cases",
			args: args{[]int{math.MaxInt, math.MinInt}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := BubbleSort(tt.args.arr); !isSorted(got) {
				t.Errorf("BubbleSort() = %v, want sorted", got)
			}
		})
	}
}

func testBubbleSortFloat(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success - sorted array",
			args: args{[]float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		},
		{
			name: "success - unsorted array",
			args: args{generateRandomSliceFloat(5)},
		},
		{
			name: "float edge cases",
			args: args{[]float64{math.MaxFloat64, 0, -math.MaxFloat64}},
		},
		{
			name: "float duplicates",
			args: args{[]float64{3.14, 1.41, 3.14, 2.71}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := BubbleSort(tt.args.arr); !isSorted(got) {
				t.Errorf("BubbleSort() = %v, want sorted", got)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	t.Run("integer slices", func(t *testing.T) {
		t.Parallel()
		testBubbleSortInt(t)
	})

	t.Run("float slices", func(t *testing.T) {
		t.Parallel()
		testBubbleSortFloat(t)
	})
}

func BenchmarkBubbleSort(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"Small", 10},
		{"Medium", 100},
		{"Large", 1000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			input := generateRandomSliceInt(bm.size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				BubbleSort(input)
			}
		})
	}
}

func BenchmarkBubbleSortFloat(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"Small", 10},
		{"Medium", 100},
		{"Large", 1000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			input := generateRandomSliceFloat(bm.size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				BubbleSort(input)
			}
		})
	}
}
