package slice

import (
	"reflect"
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
