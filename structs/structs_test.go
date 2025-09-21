package structs

import (
	"reflect"
	"testing"
	"time"
)

type Test struct {
	Name    string `updatable:"name"`
	Age     int    `updatable:"age"`
	IsAdult bool   `updatable:"is_adult"`
}

type ComplexTest struct {
	Data      Test      `updatable:"data"`
	UpdatedOn time.Time `updatable:"updated_on"`
	History   []string  `updatable:"history"`
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
				{FieldName: "name", OldValue: "example", NewValue: "example - updated"},
				{FieldName: "age", OldValue: 10, NewValue: 18},
				{FieldName: "is_adult", OldValue: false, NewValue: true},
			},
			wantErr: false,
		},
		{
			name: "success - compare two complex structs",
			old:  ComplexTest{Data: Test{Name: "example1", Age: 25, IsAdult: true}, UpdatedOn: time.Date(2024, 10, 22, 00, 00, 00, 00, time.UTC), History: []string{"user1", "user2"}},
			new:  ComplexTest{Data: Test{Name: "example1", Age: 26, IsAdult: true}, UpdatedOn: time.Date(2024, 10, 22, 00, 01, 00, 00, time.UTC), History: []string{"user1", "user2", "user3"}},
			want: []Result{
				{FieldName: "data.age", OldValue: 25, NewValue: 26},
				{FieldName: "updated_on", OldValue: time.Date(2024, 10, 22, 0, 0, 0, 0, time.UTC), NewValue: time.Date(2024, 10, 22, 0, 1, 0, 0, time.UTC)},
				{FieldName: "history", OldValue: []string{"user1", "user2"}, NewValue: []string{"user1", "user2", "user3"}},
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

type PointerTest struct {
	Name         *string            `updatable:"name_pointer"`
	Age          *int               `updatable:"age_pointer"`
	IsAdult      *bool              `updatable:"is_adult_pointer"`
	PointerSlice *[]int             `updatable:"Slice_Pointer"`
	MyMap        *map[string]string `updatable:"map"`
}

type ComplexPointerTest struct {
	Data      PointerTest       `updatable:"data"`
	UpdatedOn time.Time         `updatable:"updated_on"`
	History   []string          `updatable:"history"`
	MyMap     map[string]string `updatable:"map"`
}

func TestCompareStructsWithPointers(t *testing.T) {
	name1 := "example"
	name2 := "example2"
	age1 := 25
	age2 := 26
	isAdult := true
	slice1 := []int{1, 2, 3}
	slice2 := []int{2, 3, 4}
	map1 := map[string]string{"a": "b"}
	map2 := map[string]string{"a": "c"}

	old := ComplexPointerTest{
		Data: PointerTest{
			Name:         &name1,
			Age:          &age1,
			IsAdult:      &isAdult,
			PointerSlice: &slice1,
			MyMap:        &map1,
		},
		UpdatedOn: time.Date(2024, 10, 22, 0, 0, 0, 0, time.UTC),
		History:   []string{"user1", "user2"},
		MyMap:     map[string]string{"k": "v1"},
	}

	newVal := ComplexPointerTest{
		Data: PointerTest{
			Name:         &name2,
			Age:          &age2,
			IsAdult:      &isAdult,
			PointerSlice: &slice2,
			MyMap:        &map2,
		},
		UpdatedOn: time.Date(2024, 10, 22, 0, 1, 0, 0, time.UTC),
		History:   []string{"user1", "user2", "user3"},
		MyMap:     map[string]string{"k": "v2"},
	}

	want := []Result{
		{FieldName: "data.name_pointer", OldValue: "example", NewValue: "example2"},
		{FieldName: "data.age_pointer", OldValue: 25, NewValue: 26},
		{FieldName: "data.Slice_Pointer", OldValue: []int{1, 2, 3}, NewValue: []int{2, 3, 4}},
		{FieldName: "data.map", OldValue: map[string]string{"a": "b"}, NewValue: map[string]string{"a": "c"}},
		{FieldName: "updated_on", OldValue: time.Date(2024, 10, 22, 0, 0, 0, 0, time.UTC), NewValue: time.Date(2024, 10, 22, 0, 1, 0, 0, time.UTC)},
		{FieldName: "history", OldValue: []string{"user1", "user2"}, NewValue: []string{"user1", "user2", "user3"}},
		{FieldName: "map", OldValue: map[string]string{"k": "v1"}, NewValue: map[string]string{"k": "v2"}},
	}

	got, err := CompareStructs(old, newVal)
	if err != nil {
		t.Fatalf("CompareStructs() error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("CompareStructs() = %v\nWant: %v", got, want)
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkCompareStructsSimple(b *testing.B) {
	oldVal := Test{Name: "example", Age: 10, IsAdult: false}
	newVal := Test{Name: "example - updated", Age: 18, IsAdult: true}

	b.ReportAllocs()

	for b.Loop() {
		_, err := CompareStructs(oldVal, newVal)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompareStructsComplex(b *testing.B) {
	oldVal := ComplexTest{
		Data:      Test{Name: "example1", Age: 25, IsAdult: true},
		UpdatedOn: time.Date(2024, 10, 22, 0, 0, 0, 0, time.UTC),
		History:   []string{"user1", "user2"},
	}
	newVal := ComplexTest{
		Data:      Test{Name: "example1", Age: 26, IsAdult: true},
		UpdatedOn: time.Date(2024, 10, 22, 0, 1, 0, 0, time.UTC),
		History:   []string{"user1", "user2", "user3"},
	}

	b.ReportAllocs()

	for b.Loop() {
		_, err := CompareStructs(oldVal, newVal)
		if err != nil {
			b.Fatal(err)
		}
	}
}
