package structs

import (
	"reflect"
	"testing"
	"time"
)

type Test struct {
	Name    string `updateable:"name"`
	Age     int    `updateable:"age"`
	IsAdult bool   `updateable:"is_adult"`
}

type ComplexTest struct {
	Data      Test      `updateable:"data"`
	UpdatedOn time.Time `updateable:"updated_on"`
	History   []string  `updateable:"history"`
}

func TestCompareStructs(t *testing.T) {
	tests := []struct {
		name    string
		old     interface{}
		new     interface{}
		want    []Result
		wantErr bool
	}{
		{
			name: "success - compare two structs",
			old:  Test{Name: "example", Age: 10, IsAdult: false},
			new:  Test{Name: "example - updated", Age: 18, IsAdult: true},
			want: []Result{
				{
					FieldName: "name",
					OldValue:  "example",
					NewValue:  "example - updated",
				},
				{
					FieldName: "age",
					OldValue:  10,
					NewValue:  18,
				},
				{
					FieldName: "is_adult",
					OldValue:  false,
					NewValue:  true,
				},
			},
			wantErr: false,
		},
		{
			name: "success - compare two complex structs",
			old:  ComplexTest{Data: Test{Name: "example1", Age: 25, IsAdult: true}, UpdatedOn: time.Date(2024, 10, 22, 00, 00, 00, 00, time.UTC), History: []string{"user1", "user2"}},
			new:  ComplexTest{Data: Test{Name: "example1", Age: 26, IsAdult: true}, UpdatedOn: time.Date(2024, 10, 22, 00, 01, 00, 00, time.UTC), History: []string{"user1", "user2", "user3"}},
			want: []Result{
				{
					FieldName: "data",
					OldValue:  Test{Name: "example1", Age: 25, IsAdult: true},
					NewValue:  Test{Name: "example1", Age: 26, IsAdult: true},
				},
				{
					FieldName: "updated_on",
					OldValue:  time.Date(2024, 10, 22, 00, 00, 00, 00, time.UTC),
					NewValue:  time.Date(2024, 10, 22, 00, 01, 00, 00, time.UTC),
				},
				{
					FieldName: "history",
					OldValue:  []string{"user1", "user2"},
					NewValue:  []string{"user1", "user2", "user3"},
				},
			},
			wantErr: false,
		},
		{
			name:    "fail - non struct parameters",
			old:     map[string]string{"test": "example"},
			new:     map[string]string{"test": "example-updated"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareStructs(tt.old, tt.new)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareStructs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareStructs() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkCompareStructsSimple(b *testing.B) {
	old := Test{Name: "example", Age: 10, IsAdult: false}
	new := Test{Name: "example - updated", Age: 18, IsAdult: true}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := CompareStructs(old, new)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompareStructsComplex(b *testing.B) {
	old := ComplexTest{
		Data:      Test{Name: "example1", Age: 25, IsAdult: true},
		UpdatedOn: time.Date(2024, 10, 22, 0, 0, 0, 0, time.UTC),
		History:   []string{"user1", "user2"},
	}
	new := ComplexTest{
		Data:      Test{Name: "example1", Age: 26, IsAdult: true},
		UpdatedOn: time.Date(2024, 10, 22, 0, 1, 0, 0, time.UTC),
		History:   []string{"user1", "user2", "user3"},
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := CompareStructs(old, new)
		if err != nil {
			b.Fatal(err)
		}
	}
}
