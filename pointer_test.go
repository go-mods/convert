package convert

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToPtr tests the ToPtr function
func TestToPtr(t *testing.T) {

	t.Run("int value", func(t *testing.T) {
		value := 123
		result := ToPtr(value)

		assert.NotNil(t, result)
		assert.Exactly(t, &value, result)
		assert.Exactly(t, value, *result)
	})

	t.Run("pointer to int", func(t *testing.T) {
		value := 123
		ptr := &value
		result := ToPtr(ptr)

		assert.NotNil(t, result)
		assert.Exactly(t, &ptr, result)
		assert.Exactly(t, ptr, *result)
	})

	t.Run("string value", func(t *testing.T) {
		value := "test"
		result := ToPtr(value)

		assert.NotNil(t, result)
		assert.Exactly(t, &value, result)
		assert.Exactly(t, value, *result)
	})

	t.Run("pointer to string", func(t *testing.T) {
		value := "test"
		ptr := &value
		result := ToPtr(ptr)

		assert.NotNil(t, result)
		assert.Exactly(t, &ptr, result)
		assert.Exactly(t, ptr, *result)
	})
}

// TestIndirect tests the Indirect function
func TestIndirect(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "Non-pointer value",
			input:    42,
			expected: 42,
		},
		{
			name:     "Pointer to int",
			input:    ToPtr(123),
			expected: 123,
		},
		{
			name:     "Pointer to pointer to int",
			input:    ToPtr(ToPtr(456)),
			expected: 456,
		},
		{
			name:     "Non-pointer string",
			input:    "test",
			expected: "test",
		},
		{
			name:     "Pointer to string",
			input:    ToPtr("test"),
			expected: "test",
		},
		{
			name:     "Pointer to pointer to string",
			input:    ToPtr(ToPtr("test")),
			expected: "test",
		},
		{
			name:     "Nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "Nil int pointer",
			input:    (*int)(nil),
			expected: (*int)(nil),
		},
		{
			name:     "Nil string pointer",
			input:    (*string)(nil),
			expected: (*string)(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Indirect(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Indirect() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestToPtrBool tests the ToPtr function with boolean values
func TestToPtrBool(t *testing.T) {
	t.Run("True value", func(t *testing.T) {
		value := true
		result := ToPtr(value)

		assert.NotNil(t, result, "Result should not be nil")
		assert.Equal(t, &value, result, "Result address should be equal to the original value address")
		assert.Equal(t, value, *result, "Pointed value should be equal to the original value")
	})

	t.Run("False value", func(t *testing.T) {
		value := false
		result := ToPtr(value)

		assert.NotNil(t, result, "Result should not be nil")
		assert.Equal(t, &value, result, "Result address should be equal to the original value address")
		assert.Equal(t, value, *result, "Pointed value should be equal to the original value")
	})
}

// TestToPtrString tests the ToPtr function with string values
func TestToPtrString(t *testing.T) {
	t.Run("Non-empty string", func(t *testing.T) {
		value := "test"
		result := ToPtr(value)

		assert.NotNil(t, result, "Result should not be nil")
		assert.Equal(t, &value, result, "Result address should be equal to the original value address")
		assert.Equal(t, value, *result, "Pointed value should be equal to the original value")
	})

	t.Run("Empty string", func(t *testing.T) {
		value := ""
		result := ToPtr(value)

		assert.NotNil(t, result, "Result should not be nil")
		assert.Equal(t, &value, result, "Result address should be equal to the original value address")
		assert.Equal(t, value, *result, "Pointed value should be equal to the original value")
	})
}
