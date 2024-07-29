package convert

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToStringE(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
		hasError bool
	}{
		{"String", "test", "test", false},
		{"Bytes", []byte("test"), "test", false},
		{"Nil", nil, "", false},
		{"Bool true", true, "true", false},
		{"Bool false", false, "false", false},
		{"Int", 123, "123", false},
		{"Int8", int8(8), "8", false},
		{"Int16", int16(16), "16", false},
		{"Int32", int32(32), "32", false},
		{"Int64", int64(64), "64", false},
		{"Int64", int64(9223372036854775807), "9223372036854775807", false},
		{"Uint", uint(42), "42", false},
		{"Uint8", uint8(8), "8", false},
		{"Uint16", uint16(16), "16", false},
		{"Uint32", uint32(32), "32", false},
		{"Uint64", uint64(64), "64", false},
		{"Uint64", uint64(18446744073709551615), "18446744073709551615", false},
		{"Float32", float32(3.14), "3.14", false},
		{"Float64", 3.14159265359, "3.14159265359", false},
		{"Stringer", stringerType{"test"}, "test", false},
		{"GoStringer", goStringerTest{"test"}, "GoString: test", false},
		{"Error", errors.New("test error"), "test error", false},
		{"Complex", complex(1, 2), "(1+2i)", false},
		{"Slice", []int{1, 2, 3}, "[1 2 3]", false},
		{"Map", map[string]int{"a": 1, "b": 2}, "map[a:1 b:2]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToStringE(tt.input)
			if (err != nil) != tt.hasError {
				t.Errorf("ToStringE() error = %v, hasError %v", err, tt.hasError)
				return
			}
			if result != tt.expected {
				t.Errorf("ToStringE() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToStringWithJSONNumber(t *testing.T) {
	// Créer un StringConvert personnalisé pour json.Number
	jsonNumberConverter := func(value interface{}) *string {
		if num, ok := value.(json.Number); ok {
			str := string(num)
			return &str
		}
		return nil
	}

	// Cas de test
	testCases := []struct {
		name     string
		input    json.Number
		expected string
	}{
		{
			name:     "Nombre entier",
			input:    json.Number("42"),
			expected: "42",
		},
		{
			name:     "Nombre décimal",
			input:    json.Number("3.14"),
			expected: "3.14",
		},
		{
			name:     "Grand nombre",
			input:    json.Number("9007199254740991"),
			expected: "9007199254740991",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToString(tc.input, jsonNumberConverter)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestToStringArray(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{
			name:     "Slice of ints",
			input:    []int{1, 2, 3},
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "Array of strings",
			input:    [3]string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Single string",
			input:    "test",
			expected: []string{"test"},
		},
		{
			name:     "Slice of interface{}",
			input:    []interface{}{1, "two", 3.0},
			expected: []string{"1", "two", "3"},
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: []string{},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToStringArray(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToStringArrayWithCustomConverter(t *testing.T) {
	type CustomType struct {
		Value int
	}

	customConverter := func(value interface{}) *string {
		if v, ok := value.(CustomType); ok {
			s := fmt.Sprintf("Custom(%d)", v.Value)
			return &s
		}
		return nil
	}

	input := []CustomType{{Value: 1}, {Value: 2}, {Value: 3}}
	expected := []string{"Custom(1)", "Custom(2)", "Custom(3)"}

	result := ToStringArray(input, customConverter)
	assert.Equal(t, expected, result)
}

func TestToStringArrayWithTime(t *testing.T) {
	timeConverter := func(value interface{}) *string {
		if v, ok := value.(time.Time); ok {
			s := v.Format("2006-01-02")
			return &s
		}
		return nil
	}

	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	input := []time.Time{now, tomorrow}
	expected := []string{now.Format("2006-01-02"), tomorrow.Format("2006-01-02")}

	result := ToStringArray(input, timeConverter)
	assert.Equal(t, expected, result)
}

type stringerType struct {
	s string
}

func (st stringerType) String() string {
	return st.s
}

type goStringerTest struct {
	s string
}

func (gst goStringerTest) GoString() string {
	return fmt.Sprintf("GoString: %s", gst.s)
}
