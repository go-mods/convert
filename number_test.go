package convert

import (
	"testing"
)

func TestToIntE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
		hasError bool
	}{
		{"123", 123, false},
		{"", 0, false},
		{"abc", 0, true},
		{123, 123, false},
		{123.45, 123, false},
		{true, 1, false},
		{false, 0, false},
	}

	for _, test := range tests {
		result, err := ToIntE(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ToIntE(%v) error = %v, expected error = %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("ToIntE(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
	}{
		{"123", 123},
		{"", 0},
		{"abc", 0},
		{123, 123},
		{123.45, 123},
		{true, 1},
		{false, 0},
	}

	for _, test := range tests {
		result := ToInt(test.input)
		if result != test.expected {
			t.Errorf("ToInt(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToIntOrDefault(t *testing.T) {
	tests := []struct {
		input        interface{}
		defaultValue int
		expected     int
	}{
		{"123", 0, 123},
		{"", 0, 0},
		{"abc", 0, 0},
		{123, 0, 123},
		{123.45, 0, 123},
		{true, 0, 1},
		{false, 0, 0},
		{nil, 42, 42},
	}

	for _, test := range tests {
		result := ToIntOrDefault(test.input, test.defaultValue)
		if result != test.expected {
			t.Errorf("ToIntOrDefault(%v, %v) = %v, expected %v", test.input, test.defaultValue, result, test.expected)
		}
	}
}

func TestToUintE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint
		hasError bool
	}{
		{"123", 123, false},
		{"", 0, false},
		{"abc", 0, true},
		{123, 123, false},
		{123.45, 123, false},
		{true, 1, false},
		{false, 0, false},
	}

	for _, test := range tests {
		result, err := ToUintE(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ToUintE(%v) error = %v, expected error = %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("ToUintE(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToUint(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint
	}{
		{"123", 123},
		{"", 0},
		{"abc", 0},
		{123, 123},
		{123.45, 123},
		{true, 1},
		{false, 0},
	}

	for _, test := range tests {
		result := ToUint(test.input)
		if result != test.expected {
			t.Errorf("ToUint(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToUintOrDefault(t *testing.T) {
	tests := []struct {
		input        interface{}
		defaultValue uint
		expected     uint
	}{
		{"123", 0, 123},
		{"", 0, 0},
		{"abc", 0, 0},
		{123, 0, 123},
		{123.45, 0, 123},
		{true, 0, 1},
		{false, 0, 0},
		{nil, 42, 42},
	}

	for _, test := range tests {
		result := ToUintOrDefault(test.input, test.defaultValue)
		if result != test.expected {
			t.Errorf("ToUintOrDefault(%v, %v) = %v, expected %v", test.input, test.defaultValue, result, test.expected)
		}
	}
}

func TestToFloat32E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float32
		hasError bool
	}{
		{"123.45", 123.45, false},
		{"", 0, false},
		{"abc", 0, true},
		{123, 123, false},
		{123.45, 123.45, false},
		{true, 1, false},
		{false, 0, false},
	}

	for _, test := range tests {
		result, err := ToFloat32E(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ToFloat32E(%v) error = %v, expected error = %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("ToFloat32E(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToFloat32(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float32
	}{
		{"123.45", 123.45},
		{"", 0},
		{"abc", 0},
		{123, 123},
		{123.45, 123.45},
		{true, 1},
		{false, 0},
	}

	for _, test := range tests {
		result := ToFloat32(test.input)
		if result != test.expected {
			t.Errorf("ToFloat32(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToFloat32OrDefault(t *testing.T) {
	tests := []struct {
		input        interface{}
		defaultValue float32
		expected     float32
	}{
		{"123.45", 0, 123.45},
		{"", 0, 0},
		{"abc", 0, 0},
		{123, 0, 123},
		{123.45, 0, 123.45},
		{true, 0, 1},
		{false, 0, 0},
		{nil, 42, 42},
	}

	for _, test := range tests {
		result := ToFloat32OrDefault(test.input, test.defaultValue)
		if result != test.expected {
			t.Errorf("ToFloat32OrDefault(%v, %v) = %v, expected %v", test.input, test.defaultValue, result, test.expected)
		}
	}
}

func TestToFloat64E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
		hasError bool
	}{
		{"123.45", 123.45, false},
		{"", 0, false},
		{"abc", 0, true},
		{123, 123, false},
		{123.45, 123.45, false},
		{true, 1, false},
		{false, 0, false},
	}

	for _, test := range tests {
		result, err := ToFloat64E(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("ToFloat64E(%v) error = %v, expected error = %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("ToFloat64E(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
	}{
		{"123.45", 123.45},
		{"", 0},
		{"abc", 0},
		{123, 123},
		{123.45, 123.45},
		{true, 1},
		{false, 0},
	}

	for _, test := range tests {
		result := ToFloat64(test.input)
		if result != test.expected {
			t.Errorf("ToFloat64(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToFloat64OrDefault(t *testing.T) {
	tests := []struct {
		input        interface{}
		defaultValue float64
		expected     float64
	}{
		{"123.45", 0, 123.45},
		{"", 0, 0},
		{"abc", 0, 0},
		{123, 0, 123},
		{123.45, 0, 123.45},
		{true, 0, 1},
		{false, 0, 0},
		{nil, 42, 42},
	}

	for _, test := range tests {
		result := ToFloat64OrDefault(test.input, test.defaultValue)
		if result != test.expected {
			t.Errorf("ToFloat64OrDefault(%v, %v) = %v, expected %v", test.input, test.defaultValue, result, test.expected)
		}
	}
}
