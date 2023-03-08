package convert

// ToPtr returns a pointer to the given value
func ToPtr[T any](v T) *T {
	return &v
}

// FromPtr returns the value of the given pointer
func FromPtr[T any](v *T) T {
	return *v
}

// FromPtrCb returns the value of the given pointer or the default value if the pointer is nil
func FromPtrCb[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}
