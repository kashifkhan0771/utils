// Package errutils provides utilities for error aggregation and handling.
package errutils

import (
	"errors"
	"reflect"
	"testing"
)

func TestErrorAggregator_Add(t *testing.T) {
	// create one instance on top to use in test
	e := NewErrorAggregator()

	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		total int
	}{
		{
			name:  "success - add a new error",
			args:  args{err: errors.New("a new error")},
			total: 1,
		},
		{
			name:  "success - add a another new error",
			args:  args{err: errors.New("a new different error")},
			total: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e.Add(tt.args.err)

			if len(e.ErrorList()) != tt.total {
				t.Errorf("add error failed: got total: %d - want total: %d", len(e.ErrorList()), tt.total)
			}
		})
	}
}

func TestErrorAggregator_HasErrors(t *testing.T) {
	e := NewErrorAggregator()

	e1 := NewErrorAggregator()
	e1.Add(errors.New("a new error for test"))

	type args struct {
		errs *ErrorAggregator
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no error",
			args: args{errs: e},
			want: false,
		},
		{
			name: "has error",
			args: args{errs: e1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.errs.HasErrors(); got != tt.want {
				t.Errorf("ErrorAggregator.HasErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorAggregator_Error(t *testing.T) {
	e := NewErrorAggregator()
	e.Add(errors.New("a new error"))
	e.Add(errors.New("a new another error"))

	tests := []struct {
		name string
		errs *ErrorAggregator
		want string
	}{
		{
			name: "two errors",
			errs: e,
			want: "a new error; a new another error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// calling the Error() method to get the aggregated error
			if err := tt.errs.Error(); err != nil {
				// Comparing the error message
				if err.Error() != tt.want {
					t.Errorf("ErrorAggregator.Error() = %v, want %v", err.Error(), tt.want)
				}
			} else {
				t.Errorf("ErrorAggregator.Error() = nil, want %v", tt.want)
			}
		})
	}
}

func TestErrorAggregator_ErrorList(t *testing.T) {
	e := NewErrorAggregator()
	e.Add(errors.New("a new error"))
	e.Add(errors.New("a new another error"))

	tests := []struct {
		name string
		want []error
	}{
		{
			name: "two errors",
			want: []error{errors.New("a new error"), errors.New("a new another error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := e.ErrorList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorAggregator.ErrorList() = %v, want %v", got, tt.want)
			}
		})
	}
}
