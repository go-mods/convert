package convert

import "testing"

func TestToBool(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{"true", true},
		{"false", false},
		{1, true},
		{0, false},
		{nil, false},
		{true, true},
		{false, false},
	}

	for _, test := range tests {
		result := ToBool(test.input)
		if result != test.expected {
			t.Errorf("ToBool(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestToBoolOrDefault(t *testing.T) {
	tests := []struct {
		input        interface{}
		defaultValue bool
		expected     bool
	}{
		{"true", false, true},
		{"false", true, false},
		{1, false, true},
		{0, true, false},
		{nil, true, true},
		{nil, false, false},
		{true, false, true},
		{false, true, false},
	}

	for _, test := range tests {
		result := ToBoolOrDefault(test.input, test.defaultValue)
		if result != test.expected {
			t.Errorf("ToBoolOrDefault(%v, %v) = %v; want %v", test.input, test.defaultValue, result, test.expected)
		}
	}
}

func TestToBoolE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
		hasError bool
	}{
		{"true", true, false},
		{"false", false, false},
		{1, true, false},
		{0, false, false},
		{nil, false, false},
		{true, true, false},
		{false, false, false},
		{"invalid", false, true},
	}

	for _, test := range tests {
		result, err := ToBoolE(test.input)
		if result != test.expected || (err != nil) != test.hasError {
			t.Errorf("ToBoolE(%v) = %v, %v; want %v, %v", test.input, result, err, test.expected, test.hasError)
		}
	}
}

func TestCustomBoolConverter(t *testing.T) {
	customConverter := func(value interface{}) *bool {
		if v, ok := value.(string); ok && v == "custom" {
			b := true
			return &b
		}
		return nil
	}

	tests := []struct {
		input    interface{}
		expected bool
	}{
		{"custom", true},
		{"true", true},
		{"false", false},
		{1, true},
		{0, false},
	}

	for _, test := range tests {
		result := ToBool(test.input, customConverter)
		if result != test.expected {
			t.Errorf("ToBool(%v, customConverter) = %v; want %v", test.input, result, test.expected)
		}
	}
}
