package convert

import "reflect"

// Indirect dereferences the given value if it's a pointer.
// If v is not a pointer, it returns v as is.
// If v is a nil pointer, it returns a nil value of the pointer's type.
// For nested pointers, it recursively dereferences until a non-pointer value is found.
func Indirect(v any) any {
	// If v is nil, return a nil value
	if v == nil {
		return nil
	}

	// If v is not a pointer, return v
	if tv := reflect.TypeOf(v); tv.Kind() != reflect.Pointer {
		return v
	}

	// Get the value of v
	rv := reflect.ValueOf(v)

	// If v is nil, return a nil value of type T
	if rv.IsNil() {
		return rv.Interface()
	}

	// If v is a pointer, return the value it points to
	return Indirect(rv.Elem().Interface())
}

// ToPtr creates and returns a pointer to a copy of the given value.
// This function is useful for converting a value of any type T to a pointer to T.
func ToPtr[T any](v T) *T {
	return &v
}
