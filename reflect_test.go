package convert

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToValue(t *testing.T) {

	type args struct {
		value interface{}
		to    reflect.Type
	}

	tests := []struct {
		name    string
		args    args
		result  interface{}
		wantErr bool
	}{

		{"StringToString", args{"string", stringType}, "string", false},
		{"StringToBoolTrue", args{"1", boolType}, true, false},
		{"StringToBoolFalse", args{"0", boolType}, false, false},
		{"StringToInt", args{"-1", intType}, -1, false},
		{"StringToInt8", args{"-1", int8Type}, int8(-1), false},
		{"StringToInt16", args{"-1", int16Type}, int16(-1), false},
		{"StringToInt32", args{"-1", int32Type}, int32(-1), false},
		{"StringToInt64", args{"-1", int64Type}, int64(-1), false},
		{"StringToUint", args{"1", uintType}, uint(1), false},
		{"StringToUint8", args{"1", uint8Type}, uint8(1), false},
		{"StringToUint16", args{"1", uint16Type}, uint16(1), false},
		{"StringToUint32", args{"1", uint32Type}, uint32(1), false},
		{"StringToUint64", args{"1", uint64Type}, uint64(1), false},
		{"StringToFloat32", args{"3.14", float32Type}, float32(3.14), false},
		{"StringToFloat64", args{"3.14", float64Type}, float64(3.14), false},
		{"StringToTime", args{nowStr, timeType}, now, false},
		{"TimeToString", args{now, stringType}, nowStr, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			type TestStruct struct{ Result interface{} }
			testStruct := &TestStruct{}

			structValue := reflect.ValueOf(testStruct).Elem()
			structFieldValue := structValue.FieldByName("Result")

			value, err := ToValueE(tt.args.value, tt.args.to)
			assert.NoError(t, err)

			structFieldValue.Set(value)

			if tt.args.to == timeType {
				assert.WithinDuration(t, tt.result.(time.Time), testStruct.Result.(time.Time), time.Second)
			} else {
				assert.Exactly(t, tt.result, testStruct.Result)
			}
		})
	}
}

func TestToJsonValue(t *testing.T) {

	type Person struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		FullName  string `json:"full_name"`
	}
	var person = &Person{FirstName: "first", LastName: "last", FullName: "full"}
	var personStr = "{\"first_name\":\"first\",\"last_name\":\"last\",\"full_name\":\"full\"}"

	value, err := ToJsonValueE(personStr, reflect.TypeOf(*person))
	assert.NoError(t, err)

	personFromValue := value.Interface().(Person)
	assert.Exactly(t, *person, personFromValue)
}

func TestGetConvertType(t *testing.T) {
	now := time.Now()
	duration := time.Hour

	tests := []struct {
		name     string
		value    interface{}
		expected reflect.Type
	}{
		// Tests signed integers
		{
			name:     "integer 1",
			value:    1,
			expected: int64Type,
		},
		{
			name:     "int8",
			value:    int8(1),
			expected: int64Type,
		},
		{
			name:     "int16",
			value:    int16(1),
			expected: int64Type,
		},
		{
			name:     "int32",
			value:    int32(1),
			expected: int64Type,
		},
		{
			name:     "int64",
			value:    int64(1),
			expected: int64Type,
		},
		// Tests unsigned integers
		{
			name:     "uint",
			value:    uint(1),
			expected: uint64Type,
		},
		{
			name:     "uint8",
			value:    uint8(1),
			expected: uint64Type,
		},
		{
			name:     "uint16",
			value:    uint16(1),
			expected: uint64Type,
		},
		{
			name:     "uint32",
			value:    uint32(1),
			expected: uint64Type,
		},
		{
			name:     "uint64",
			value:    uint64(1),
			expected: uint64Type,
		},
		// Tests floating-point numbers
		{
			name:     "float32",
			value:    float32(3.14),
			expected: float32Type,
		},
		{
			name:     "float64",
			value:    float64(3.14),
			expected: float64Type,
		},
		{
			name:     "float64 as integer",
			value:    float64(1.0),
			expected: int64Type,
		},
		// Tests strings
		{
			name:     "string",
			value:    "test",
			expected: stringType,
		},
		{
			name:     "string numeric",
			value:    "1",
			expected: int64Type,
		},
		{
			name:     "string numeric",
			value:    "123",
			expected: int64Type,
		},
		{
			name:     "string float",
			value:    "3.14",
			expected: float64Type,
		},
		{
			name:     "string bool true",
			value:    "true",
			expected: boolType,
		},
		{
			name:     "string bool false",
			value:    "false",
			expected: boolType,
		},
		// Tests boolean values
		{
			name:     "bool true",
			value:    true,
			expected: boolType,
		},
		{
			name:     "bool false",
			value:    false,
			expected: boolType,
		},
		// Tests time values
		{
			name:     "time",
			value:    now,
			expected: timeType,
		},
		{
			name:     "duration",
			value:    duration,
			expected: durationType,
		},
		// Tests special cases
		{
			name:     "nil",
			value:    nil,
			expected: nilType,
		},
		{
			name:     "string empty",
			value:    "",
			expected: stringType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input value type: %v", reflect.TypeOf(tt.value))
			got := GetConvertType(tt.value)
			if got != tt.expected {
				t.Errorf("GetConvertType() = %v, want %v", got, tt.expected)
			}
		})
	}
}
