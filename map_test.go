package convert

import (
	"reflect"
	"testing"
)

func TestToMapStringString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]string
	}{
		{
			name:     "Map string string",
			input:    map[string]string{"key": "value"},
			expected: map[string]string{"key": "value"},
		},
		{
			name:     "Map interface interface",
			input:    map[interface{}]interface{}{"key": "value"},
			expected: map[string]string{"key": "value"},
		},
		{
			name:     "JSON string",
			input:    `{"key": "value"}`,
			expected: map[string]string{"key": "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMapStringString(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ToMapStringString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToMapStringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]bool
	}{
		{
			name:     "Map string bool",
			input:    map[string]bool{"key": true},
			expected: map[string]bool{"key": true},
		},
		{
			name:     "Map string interface",
			input:    map[string]interface{}{"key": true},
			expected: map[string]bool{"key": true},
		},
		{
			name:     "JSON string",
			input:    `{"key": true}`,
			expected: map[string]bool{"key": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMapStringBool(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ToMapStringBool() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToMapStringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]int
	}{
		{
			name:     "Map string int",
			input:    map[string]int{"key": 42},
			expected: map[string]int{"key": 42},
		},
		{
			name:     "Map interface interface",
			input:    map[interface{}]interface{}{"key": 42},
			expected: map[string]int{"key": 42},
		},
		{
			name:     "JSON string",
			input:    `{"key": 42}`,
			expected: map[string]int{"key": 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMapStringInt(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ToMapStringInt() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToMapStringFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]float64
	}{
		{
			name:     "Map string float64",
			input:    map[string]float64{"key": 3.14},
			expected: map[string]float64{"key": 3.14},
		},
		{
			name:     "Map string interface",
			input:    map[string]interface{}{"key": 3.14},
			expected: map[string]float64{"key": 3.14},
		},
		{
			name:     "JSON string",
			input:    `{"key": 3.14}`,
			expected: map[string]float64{"key": 3.14},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMapStringFloat64(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ToMapStringFloat64() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToMapStringInterface(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
	}{
		{
			name:     "Map string interface",
			input:    map[string]interface{}{"key": "value"},
			expected: map[string]interface{}{"key": "value"},
		},
		{
			name:     "Map interface interface",
			input:    map[interface{}]interface{}{"key": "value"},
			expected: map[string]interface{}{"key": "value"},
		},
		{
			name:     "JSON string",
			input:    `{"key": "value"}`,
			expected: map[string]interface{}{"key": "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMapStringInterface(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ToMapStringInterface() = %v, want %v", result, tt.expected)
			}
		})
	}
}
