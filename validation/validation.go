// Package validation provides validators for common checks such as strings,
// numbers, emails, URLs and slugs, plus error aggregation and struct validation.
package validation

import (
	"fmt"
	"net/mail"
	"regexp"
	"slices"
	"strings"

	"github.com/kashifkhan0771/utils/errutils"
	"github.com/kashifkhan0771/utils/url"
)

// StringNotEmpty checks that a string is not empty.
func StringNotEmpty(name, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s: must not be empty", name)
	}

	return nil
}

// StringIn checks if a value is in the allowed list.
func StringIn(name string, value string, allowed ...string) error {
	if len(allowed) == 0 {
		return nil
	}
	if slices.Contains(allowed, value) {
		return nil
	}

	return fmt.Errorf("%s: must be one of [%s]", name, strings.Join(allowed, ", "))
}

// IntBetween checks if an integer is between a range min and max.
func IntBetween(name string, value, min, max int) error {
	if value >= min && value <= max {
		return nil
	}

	return fmt.Errorf("%s: integer must be between %d and %d", name, min, max)
}

// IsEmail checks if a string is a valid email address.
func IsEmail(name, email string) error {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	if addr.Address != email {
		return fmt.Errorf("%s: email must be valid", name)
	}

	return nil
}

// IsURL checks if a string is a valid URL.
func IsURL(name, value string, allowedSchemes ...string) error {
	if len(allowedSchemes) == 0 {
		// default schemes
		allowedSchemes = []string{"https", "http"}
	}
	valid := url.IsValidURL(value, allowedSchemes)
	if !valid {
		return fmt.Errorf("%s: must be a valid URL", name)
	}

	return nil
}

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

// IsSlug checks if a string is a valid slug.
func IsSlug(name, value string) error {
	if !slugRegex.MatchString(value) {
		return fmt.Errorf("%s: must be a valid slug", name)
	}

	return nil
}

// Validator is an interface that types implement to validate multiple values.
type Validator interface {
	Validate() error
}

// ValidateStruct validates a struct that implements the Validator interface.
func ValidateStruct(v Validator) error {
	if v == nil {
		return nil
	}

	return v.Validate()
}

// Errors is a struct that holds errors when a validator is validating.
type Errors struct {
	agg *errutils.ErrorAggregator
}

// Add appends an error to the aggregator.
func (e *Errors) Add(err error) {
	if e.agg == nil {
		e.agg = errutils.NewErrorAggregator()
	}
	e.agg.Add(err)
}

// AsError returns the aggregated errors as a simple error message.
func (e *Errors) AsError() error {
	if e.agg == nil {
		return nil
	}

	return e.agg.Error()
}
