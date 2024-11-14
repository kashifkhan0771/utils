// Package errutils provides utilities for error aggregation and handling.
package errutils

import (
	"errors"
	"strings"
)

// ErrorAggregator aggregates multiple errors into a single error.
type ErrorAggregator struct {
	errors []error
}

// NewErrorAggregator creates a new instance of ErrorAggregator.
func NewErrorAggregator() *ErrorAggregator {
	return &ErrorAggregator{}
}

// Add adds a new error to the aggregator. If err is nil, it is ignored.
func (e *ErrorAggregator) Add(err error) {
	if err != nil {
		e.errors = append(e.errors, err)
	}
}

// HasErrors returns true if there are any aggregated errors.
func (e *ErrorAggregator) HasErrors() bool {
	return len(e.errors) > 0
}

// Error returns the aggregated errors as a single error message.
// If there are no errors, it returns nil.
func (e *ErrorAggregator) Error() error {
	if !e.HasErrors() {
		return nil
	}

	var sb strings.Builder

	// loop over all the errors in the list and write a string
	for i, err := range e.ErrorList() {
		sb.WriteString(err.Error())

		// for each error add `;` in the end unless it's the last one
		if i < len(e.errors)-1 {
			sb.WriteString("; ")
		}
	}

	return errors.New(sb.String())
}

// ErrorList returns the list of aggregated errors as a slice.
func (e *ErrorAggregator) ErrorList() []error {
	result := make([]error, len(e.errors))
	copy(result, e.errors)
	return result
}
