package convert_test

import (
	"github.com/go-mods/convert"
	"github.com/golang-module/carbon/v2"
	"reflect"
	"testing"
	"time"
)

var typString = reflect.TypeOf("")
var typBool = reflect.TypeOf(false)
var typInt = reflect.TypeOf(0)
var typInt8 = reflect.TypeOf(int8(0))
var typInt16 = reflect.TypeOf(int16(0))
var typInt32 = reflect.TypeOf(int32(0))
var typInt64 = reflect.TypeOf(int64(0))
var typUint = reflect.TypeOf(uint(0))
var typUint8 = reflect.TypeOf(uint8(0))
var typUint16 = reflect.TypeOf(uint16(0))
var typUint32 = reflect.TypeOf(uint32(0))
var typUint64 = reflect.TypeOf(uint64(0))
var typFloat32 = reflect.TypeOf(float32(0))
var typFloat64 = reflect.TypeOf(float64(0))
var typTime = reflect.TypeOf(time.Now())

//var typTimeDuration = reflect.TypeOf(time.Duration(0))

var now = carbon.CreateFromTimestamp(123456789).Carbon2Time()
var nowStr = carbon.CreateFromTimestamp(123456789).ToRfc3339String()

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

		{"StringToString", args{"string", typString}, "string", false},
		{"StringToBoolTrue", args{"1", typBool}, true, false},
		{"StringToBoolFalse", args{"0", typBool}, false, false},
		{"StringToInt", args{"-1", typInt}, -1, false},
		{"StringToInt8", args{"-1", typInt8}, int8(-1), false},
		{"StringToInt16", args{"-1", typInt16}, int16(-1), false},
		{"StringToInt32", args{"-1", typInt32}, int32(-1), false},
		{"StringToInt64", args{"-1", typInt64}, int64(-1), false},
		{"StringToUint", args{"1", typUint}, uint(1), false},
		{"StringToUint8", args{"1", typUint8}, uint8(1), false},
		{"StringToUint16", args{"1", typUint16}, uint16(1), false},
		{"StringToUint32", args{"1", typUint32}, uint32(1), false},
		{"StringToUint64", args{"1", typUint64}, uint64(1), false},
		{"StringToFloat32", args{"3.14", typFloat32}, float32(3.14), false},
		{"StringToFloat64", args{"3.14", typFloat64}, float64(3.14), false},
		{"StringToTime", args{nowStr, typTime}, now, false},

		{"TimeToString", args{now, typString}, nowStr, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			type TestStruct struct{ Result interface{} }
			testStruct := &TestStruct{}

			structValue := reflect.ValueOf(testStruct).Elem()
			structFieldValue := structValue.FieldByName("Result")
			value, err := convert.ToValue(tt.args.value, tt.args.to)

			if (err != nil) != tt.wantErr {
				t.Errorf("ToValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			structFieldValue.Set(value)

			if !reflect.DeepEqual(tt.result, testStruct.Result) {
				t.Errorf("ToValue() got = %v, want %v", testStruct.Result, tt.result)
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

	value, err := convert.ToJsonValue(personStr, reflect.TypeOf(*person))
	personFromValue := value.Interface().(Person)

	if err != nil {
		t.Errorf("TestToJsonValue() error = %v", err)
		return
	}

	if *person != personFromValue {
		t.Errorf("ToValue() got = %v, want %v", personFromValue, person)
	}

}
