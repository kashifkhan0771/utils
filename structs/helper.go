package structs

import "reflect"

// derefOrZeroStruct returns a usable reflect.Value for struct comparison.
//
// This is used specifically for structs or pointers to structs.
// - If v is a non-pointer struct, it is returned as-is.
// - If v is a pointer to a struct:
//   - If the pointer is nil, it returns a zero-initialized struct.
//   - If the pointer is non-nil, it is dereferenced and returned.
//
// This allows CompareStructs to safely handle nested nil struct fields like `*Profile == nil`
func derefOrZeroStruct(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflect.Zero(v.Type().Elem())
		}

		return v.Elem()
	}

	return v
}

// derefIfPointer returns the dereferenced value if v is a non-nil pointer.
//
// This is used for primitives, slices, maps, arrays, and other non-struct types.
// - If v is a non-nil pointer, it returns the dereferenced value.
// - If v is not a pointer or is nil, it returns v as-is.
//
// Used in field-by-field comparison to avoid comparing pointers directly.
func derefIfPointer(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr && !v.IsNil() {

		return v.Elem()
	}

	return v
}
