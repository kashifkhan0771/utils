package slice

import (
	"crypto/rand"
	"math/big"
	"reflect"
	"strings"
	"testing"
)

func TestRemoveDuplicateStr(t *testing.T) {
	type args struct {
		strSlice []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success - remove duplicate strings from slice",
			args: args{[]string{"one", "one", "one", "two", "three"}},
			want: []string{"one", "two", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicateStr(tt.args.strSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicateStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicateInt(t *testing.T) {
	type args struct {
		strSlice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success - remove duplicate integer from a slice",
			args: args{[]int{1, 2, 3, 4, 4, 5, 5, 6, 7, 7, 7}},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicateInt(tt.args.strSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicateInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func generateStrings(n int) []string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/~`"
	data := make([]string, n)

	for i := range n {
		strLen := (i % 10) + 5 // Generate strings of length 5-14
		var strBuilder strings.Builder
		for j := 0; j < strLen; j++ {
			strBuilder.WriteByte(charset[(i+j)%len(charset)])
		}
		data[i] = strBuilder.String()
	}

	return data
}

func generateRandomInts(n int) []int {
	maxVal := big.NewInt(1000)
	data := make([]int, n)

	for i := range n {
		num, err := rand.Int(rand.Reader, maxVal)
		if err != nil {
			panic(err)
		}
		data[i] = int(num.Int64())
	}

	return data
}

func BenchmarkRemoveDuplicateStrings(b *testing.B) {
	data := generateStrings(100000)

	b.ReportAllocs()

	for b.Loop() {
		RemoveDuplicateStr(data)
	}
}

func BenchmarkRemoveDuplicateInts(b *testing.B) {
	data := generateRandomInts(100000)

	b.ReportAllocs()

	for b.Loop() {
		RemoveDuplicateInt(data)
	}
}
