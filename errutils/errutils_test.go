// Package errutils provides utilities for error aggregation and handling.
package errutils

import (
	"errors"
	"reflect"
	"testing"
)

func TestErrorAggregator_Add(t *testing.T) {
	t.Parallel()

	// create one instance on top to use in test
	e := NewErrorAggregator()

	tests := []struct {
		name  string
		err   error
		total int
	}{
		{
			name:  "success - add a new error",
			err:   errors.New("a new error"),
			total: 1,
		},
		{
			name:  "success - add another new error",
			err:   errors.New("a new different error"),
			total: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e.Add(tt.err)

			if len(e.ErrorList()) != tt.total {
				t.Errorf("add error failed: got total: %d - want total: %d", len(e.ErrorList()), tt.total)
			}
		})
	}
}

func TestErrorAggregator_HasErrors(t *testing.T) {
	t.Parallel()

	e := NewErrorAggregator()
	e.Add(errors.New("a new error for test"))

	tests := []struct {
		name string
		errs *ErrorAggregator
		want bool
	}{
		{
			name: "no error",
			errs: NewErrorAggregator(),
			want: false,
		},
		{
			name: "has error",
			errs: e,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.errs.HasErrors(); got != tt.want {
				t.Errorf("ErrorAggregator.HasErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorAggregator_Error(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

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
	t.Parallel()

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
			t.Parallel()

			if got := e.ErrorList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorAggregator.ErrorList() = %v, want %v", got, tt.want)
			}
		})
	}
}
