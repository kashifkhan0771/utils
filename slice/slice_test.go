package slice

import (
	"math/rand"
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

	for i := 0; i < n; i++ {
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
	r := rand.New(rand.NewSource(99))

	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = r.Intn(1000)
	}
	return data
}

func BenchmarkRemoveDuplicateStrings(b *testing.B) {
	data := generateStrings(100000)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RemoveDuplicateStr(data)
	}
}

func BenchmarkRemoveDuplicateInts(b *testing.B) {
	data := generateRandomInts(100000)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RemoveDuplicateInt(data)
	}
}
