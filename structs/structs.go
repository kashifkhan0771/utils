package structs

import (
	"fmt"
	"reflect"
	"strings"
)

// Result represents the comparison outcome of a field, including its name, old value, and new value.
type Result struct {
	FieldName string
	OldValue  interface{}
	NewValue  interface{}
}

// CompareStructs compares two struct instances (or pointers to structs) and returns their differences.
//
// Only fields explicitly tagged with `updatable` are compared. Supported tag formats:
//   - `updatable:"true"`        → uses the struct field name
//   - `updatable:"custom_name"` → uses the custom field name in results
//
// This function supports:
//   - Structs and pointers to structs (including nil pointers)
//   - Recursively comparing nested structs (excluding time.Time)
//   - Comparing primitives, slices, arrays, maps, and pointers
//
// If a field's value differs (based on reflect.DeepEqual), the field is added to the result.
//
// The optional `prefix` is used for naming nested fields like `data.age`.
func CompareStructs(old, new interface{}, prefix ...string) ([]Result, error) {
	if reflect.TypeOf(old) != reflect.TypeOf(new) {
		return nil, fmt.Errorf("both structs must be of the same type")
	}

	//Checks for nil and dereferences
	oldVal := derefOrZeroStruct(reflect.ValueOf(old))
	newVal := derefOrZeroStruct(reflect.ValueOf(new))

	if oldVal.Kind() != reflect.Struct || newVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct or pointer to struct")
	}

	var results []Result

	for i := 0; i < oldVal.NumField(); i++ {
		field := oldVal.Type().Field(i)

		if !field.IsExported() {
			continue
		}

		oldField := oldVal.Field(i)
		newField := newVal.Field(i)

		// Check if `updatable` tag exists
		if tag, ok := field.Tag.Lookup("updatable"); ok && tag != "" {
			fieldName := field.Name
			if tag != "true" {
				fieldName = tag
			}
			fullName := strings.Join(append(prefix, fieldName), ".")

			// Handle nested structs or pointers to struc
			if oldField.Kind() == reflect.Struct && oldField.Type().Name() != "Time" {
				// Recurse for nested structs (ignore time.Time)
				nestedOld := oldField.Interface()
				nestedNew := newField.Interface()

				nestedResults, err := CompareStructs(nestedOld, nestedNew, append(prefix, fieldName)...)
				if err != nil {
					return nil, err
				}
				results = append(results, nestedResults...)

				continue
			}

			// For everything else (primitive, pointer, slice, map, array, etc.)
			oldVal := derefIfPointer(oldField)
			newVal := derefIfPointer(newField)

			if !reflect.DeepEqual(oldVal.Interface(), newVal.Interface()) {
				results = append(results, Result{
					FieldName: fullName,
					OldValue:  oldVal.Interface(),
					NewValue:  newVal.Interface(),
				})
			}
		}
	}

	return results, nil
}
