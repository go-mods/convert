package convert

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
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
