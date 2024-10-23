package structs

import (
	"fmt"
	"reflect"
)

// Result represents the comparison outcome of a field, including its name, old value, and new value.
type Result struct {
	FieldName string
	OldValue  interface{}
	NewValue  interface{}
}

/*
CompareStructs compares two struct instances of the same type
and returns a list of results with the old and new values of each field tagged with `updateable`.
*/
func CompareStructs(old, new interface{}) ([]Result, error) {
	if reflect.TypeOf(old) != reflect.TypeOf(new) {
		return nil, fmt.Errorf("both structs must be of the same type")
	}

	oldValue := reflect.ValueOf(old)
	newValue := reflect.ValueOf(new)

	var comparedResults = make([]Result, 0)

	for i := 0; i < oldValue.NumField(); i++ {
		field := oldValue.Type().Field(i)

		if !field.IsExported() {
			continue // skip unexported fields
		}

		oldFieldValue := oldValue.Field(i)
		newFieldValue := newValue.Field(i)

		// check if the field has the `updateable` tag
		if tag, ok := field.Tag.Lookup("updateable"); ok && tag != "" {
			fieldName := field.Name // default to the struct field name
			if tag != "true" {      // if a custom tag is provided, use that as the field name
				fieldName = tag
			}

			if !reflect.DeepEqual(oldFieldValue.Interface(), newFieldValue.Interface()) {
				comparedResults = append(comparedResults, Result{
					FieldName: fieldName,
					OldValue:  oldFieldValue.Interface(),
					NewValue:  newFieldValue.Interface(),
				})
			}
		}

	}

	return comparedResults, nil
}
