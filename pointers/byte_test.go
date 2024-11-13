package pointers

import (
	"reflect"
	"testing"
)

func TestNullableByteSlice(t *testing.T) {
	type args struct {
		b *[]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "success - b is nil",
			args: args{
				b: nil,
			},
			want: []byte{},
		},
		{
			name: "success - b is an empty slice",
			args: args{
				b: new([]byte),
			},
			want: []byte{},
		},
		{
			name: "success - b is a non-empty slice",
			args: args{
				b: func() *[]byte { v := []byte{1, 2, 3}; return &v }(),
			},
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NullableByteSlice(tt.args.b)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NullableByteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
