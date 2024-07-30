package convert

import (
	"reflect"
	"testing"
	"time"
)

func TestToSliceStringE(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    []string
		expectError bool
	}{
		{"Slice of strings", []string{"a", "b", "c"}, []string{"a", "b", "c"}, false},
		{"Slice of interfaces", []interface{}{"a", "b", "c"}, []string{"a", "b", "c"}, false},
		{"Valid JSON string", `["a", "b", "c"]`, []string{"a", "b", "c"}, false},
		{"Invalid JSON string", `{"key": "value"}`, nil, true},
		{"Unsupported type", 123, nil, true},
		{"Nil value", nil, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToSliceStringE(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("ToSliceStringE(%v) should return an error", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("ToSliceStringE(%v) returned an unexpected error: %v", tt.input, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToSliceStringE(%v) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestToSliceInterfaceE(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    []interface{}
		expectError bool
	}{
		{"Slice of interfaces", []interface{}{1, "b", true}, []interface{}{1, "b", true}, false},
		{"Valid JSON string", `[1, "b", true]`, []interface{}{float64(1), "b", true}, false},
		{"Invalid JSON string", `{"key": "value"}`, nil, true},
		{"Unsupported type", 123, nil, true},
		{"Nil value", nil, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToSliceInterfaceE(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("ToSliceInterfaceE(%v) should return an error", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("ToSliceInterfaceE(%v) returned an unexpected error: %v", tt.input, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToSliceInterfaceE(%v) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestToSliceBoolE(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    []bool
		expectError bool
	}{
		{"Slice of bools", []bool{true, false, true}, []bool{true, false, true}, false},
		{"Slice of interfaces", []interface{}{true, false, true}, []bool{true, false, true}, false},
		{"Valid JSON string", `[true, false, true]`, []bool{true, false, true}, false},
		{"Invalid JSON string", `{"key": "value"}`, nil, true},
		{"Unsupported type", 123, nil, true},
		{"Nil value", nil, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToSliceBoolE(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("ToSliceBoolE(%v) should return an error", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("ToSliceBoolE(%v) returned an unexpected error: %v", tt.input, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToSliceBoolE(%v) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestToSliceIntE(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    []int
		expectError bool
	}{
		{"Slice of ints", []int{1, 2, 3}, []int{1, 2, 3}, false},
		{"Slice of interfaces", []interface{}{1, 2, 3}, []int{1, 2, 3}, false},
		{"Valid JSON string", `[1, 2, 3]`, []int{1, 2, 3}, false},
		{"Invalid JSON string", `{"key": "value"}`, nil, true},
		{"Unsupported type", "abc", nil, true},
		{"Nil value", nil, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToSliceIntE(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("ToSliceIntE(%v) should return an error", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("ToSliceIntE(%v) returned an unexpected error: %v", tt.input, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToSliceIntE(%v) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestToSliceTimeE(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name        string
		input       interface{}
		expected    []time.Time
		expectError bool
	}{
		{"Slice of time.Time", []time.Time{now, now.Add(time.Hour)}, []time.Time{now, now.Add(time.Hour)}, false},
		{"Slice of interfaces", []interface{}{now, now.Add(time.Hour)}, []time.Time{now, now.Add(time.Hour)}, false},
		{"Invalid JSON string", `{"key": "value"}`, nil, true},
		{"Unsupported type", 123, nil, true},
		{"Nil value", nil, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToSliceTimeE(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("ToSliceTimeE(%v) should return an error", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("ToSliceTimeE(%v) returned an unexpected error: %v", tt.input, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToSliceTimeE(%v) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestToSliceDurationE(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    []time.Duration
		expectError bool
	}{
		{"Slice of time.Duration", []time.Duration{time.Hour, time.Minute}, []time.Duration{time.Hour, time.Minute}, false},
		{"Slice of interfaces", []interface{}{time.Hour, time.Minute}, []time.Duration{time.Hour, time.Minute}, false},
		{"Invalid JSON string", `{"key": "value"}`, nil, true},
		{"Unsupported type", 123, nil, true},
		{"Nil value", nil, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToSliceDurationE(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("ToSliceDurationE(%v) should return an error", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("ToSliceDurationE(%v) returned an unexpected error: %v", tt.input, err)
				}
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ToSliceDurationE(%v) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}
