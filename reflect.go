package convert

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"
	"unicode"
)

// Caster is a function type for custom type casting.
type Caster func(value interface{}) reflect.Value

// CasterE is a function type for custom type casting with error.
type CasterE func(value interface{}) (reflect.Value, error)

// CasterConvert is an interface for custom type casting.
type CasterConvert func(value interface{}) reflect.Value

var (
	stringType   = reflect.TypeOf("")
	boolType     = reflect.TypeOf(true)
	intType      = reflect.TypeOf(int(0))
	int8Type     = reflect.TypeOf(int8(0))
	int16Type    = reflect.TypeOf(int16(0))
	int32Type    = reflect.TypeOf(int32(0))
	int64Type    = reflect.TypeOf(int64(0))
	uintType     = reflect.TypeOf(uint(0))
	uint8Type    = reflect.TypeOf(uint8(0))
	uint16Type   = reflect.TypeOf(uint16(0))
	uint32Type   = reflect.TypeOf(uint32(0))
	uint64Type   = reflect.TypeOf(uint64(0))
	float32Type  = reflect.TypeOf(float32(0))
	float64Type  = reflect.TypeOf(float64(0))
	timeType     = reflect.TypeOf((*time.Time)(nil)).Elem()
	durationType = reflect.TypeOf(time.Duration(0))
	nilType      = reflect.TypeOf(nil)
	InvalidType  = reflect.TypeOf(reflect.Value{})

	InvalidValue = reflect.Value{}
)

var casters = map[reflect.Type]CasterE{
	stringType:   castStringE,
	boolType:     castBoolE,
	intType:      castIntE,
	int8Type:     castInt8E,
	int16Type:    castInt16E,
	int32Type:    castInt32E,
	int64Type:    castInt64E,
	uintType:     castUintE,
	uint8Type:    castUint8E,
	uint16Type:   castUint16E,
	uint32Type:   castUint32E,
	uint64Type:   castUint64E,
	float32Type:  castFloat32E,
	float64Type:  castFloat64E,
	timeType:     castTimeE,
	durationType: castTimeDurationE,
}

// ToValue converts a value to a specified type using custom casters.
func ToValue(value interface{}, to reflect.Type, converters ...CasterConvert) reflect.Value {
	res, _ := ToValueE(value, to, converters...)
	return res
}

// ToValueE converts a value to a specified type using custom casters with error.
func ToValueE(value interface{}, to reflect.Type, converters ...CasterConvert) (reflect.Value, error) {
	for _, converter := range converters {
		if result := converter(value); result.IsValid() {
			return result, nil
		}
	}

	v := Indirect(value)

	if caster, ok := casters[to]; ok {
		return caster(v)
	}

	return InvalidValue, errors.New("casting not supported")
}

// ToJsonValue converts a value to a specified type using JSON unmarshalling.
func ToJsonValue(value interface{}, to reflect.Type, converters ...CasterConvert) reflect.Value {
	res, _ := ToJsonValueE(value, to, converters...)
	return res
}

// ToJsonValueE converts a value to a specified type using JSON unmarshalling with error.
func ToJsonValueE(value interface{}, to reflect.Type, converters ...CasterConvert) (reflect.Value, error) {
	for _, converter := range converters {
		if result := converter(value); result.IsValid() {
			return result, nil
		}
	}

	v := Indirect(value)

	jsonString, err := ToStringE(v)
	if err != nil {
		return InvalidValue, err
	}

	jsonValuePtr := reflect.New(to)
	jsonValue := jsonValuePtr.Elem()

	err = json.Unmarshal([]byte(jsonString), jsonValue.Addr().Interface())
	if err != nil {
		return InvalidValue, err
	}

	return jsonValue, nil
}

// func castString(value interface{}) reflect.Value {
//	res, _ := castStringE(value)
//	return res
// }

func castStringE(value interface{}) (reflect.Value, error) {
	var s string
	var err error

	if reflect.TypeOf(value) == timeType {
		s = value.(time.Time).Format(time.RFC3339)
	} else {
		s, err = ToStringE(value)
		if err != nil {
			return InvalidValue, err
		}
	}

	return reflect.ValueOf(s), nil
}

// func castBool(value interface{}) reflect.Value {
//	res, _ := castBoolE(value)
//	return res
// }

func castBoolE(value interface{}) (reflect.Value, error) {
	v, err := ToBoolE(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castInt(value interface{}) reflect.Value {
//	res, _ := castIntE(value)
//	return res
// }

func castIntE(value interface{}) (reflect.Value, error) {
	v, err := ToIntE(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castInt8(value interface{}) reflect.Value {
//	res, _ := castInt8E(value)
//	return res
// }

func castInt8E(value interface{}) (reflect.Value, error) {
	v, err := ToInt8E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castInt16(value interface{}) reflect.Value {
//	res, _ := castInt16E(value)
//	return res
// }

func castInt16E(value interface{}) (reflect.Value, error) {
	v, err := ToInt16E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castInt32(value interface{}) reflect.Value {
//	res, _ := castInt32E(value)
//	return res
// }

func castInt32E(value interface{}) (reflect.Value, error) {
	v, err := ToInt32E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castInt64(value interface{}) reflect.Value {
//	res, _ := castInt64E(value)
//	return res
// }

func castInt64E(value interface{}) (reflect.Value, error) {
	v, err := ToInt64E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castUint(value interface{}) reflect.Value {
//	res, _ := castUintE(value)
//	return res
// }

func castUintE(value interface{}) (reflect.Value, error) {
	v, err := ToUintE(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castUint8(value interface{}) reflect.Value {
//	res, _ := castUint8E(value)
//	return res
// }

func castUint8E(value interface{}) (reflect.Value, error) {
	v, err := ToUint8E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castUint16(value interface{}) reflect.Value {
//	res, _ := castUint16E(value)
//	return res
// }

func castUint16E(value interface{}) (reflect.Value, error) {
	v, err := ToUint16E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castUint32(value interface{}) reflect.Value {
//	res, _ := castUint32E(value)
//	return res
// }

func castUint32E(value interface{}) (reflect.Value, error) {
	v, err := ToUint32E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castUint64(value interface{}) reflect.Value {
//	res, _ := castUint64E(value)
//	return res
// }

func castUint64E(value interface{}) (reflect.Value, error) {
	v, err := ToUint64E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castFloat32(value interface{}) reflect.Value {
//	res, _ := castFloat32E(value)
//	return res
// }

func castFloat32E(value interface{}) (reflect.Value, error) {
	v, err := ToFloat32E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castFloat64(value interface{}) reflect.Value {
//	res, _ := castFloat64E(value)
//	return res
// }

func castFloat64E(value interface{}) (reflect.Value, error) {
	v, err := ToFloat64E(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castTime(value interface{}) reflect.Value {
//	res, _ := castTimeE(value)
//	return res
// }

func castTimeE(value interface{}) (reflect.Value, error) {
	v, err := ToTimeE(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// func castTimeDuration(value interface{}) reflect.Value {
//	res, _ := castTimeDurationE(value)
//	return res
// }

func castTimeDurationE(value interface{}) (reflect.Value, error) {
	v, err := ToDurationE(value)
	if err != nil {
		return InvalidValue, err
	}

	return reflect.ValueOf(v), nil
}

// IsAlphanumeric checks if the given string consists of only alphanumeric characters.
func IsAlphanumeric(value interface{}) bool {
	// Check if the value is a string
	str, ok := value.(string)
	if !ok {
		return false
	}

	// Iterate over each character in the string
	for _, char := range str {
		if !unicode.IsDigit(char) && unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

// GetConvertType returns the reflect.Type of a value based on its conversion.
func GetConvertType(value interface{}) reflect.Type {
	if value == nil {
		return nilType
	}

	// First check the actual type of the value
	actualType := reflect.TypeOf(value)
	actualKind := actualType.Kind()

	// If it's already a numeric type, use specific logic
	switch actualKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// Vérifier si c'est une durée
		if _, ok := value.(time.Duration); ok {
			return durationType
		}
		return int64Type // Always return int64 for integer types
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint64Type // Always return uint64 for unsigned integer types
	case reflect.Float32:
		return float32Type
	case reflect.Float64:
		// For float64, check if it's actually an integer
		f := value.(float64)
		if f == float64(int(f)) {
			return int64Type
		}
		return float64Type
	case reflect.Bool:
		return boolType
	case reflect.String:
		// Pour les chaînes vides, retourner directement stringType
		if v := value.(string); v == "" {
			return stringType
		}
	}

	// For non-numeric types or types that require conversion
	// try to find the type

	// Check special types first
	if _, err := ToTimeE(value); err == nil {
		return timeType
	}
	if _, err := ToDurationE(value); err == nil {
		return durationType
	}
	if IsAlphanumeric(value) {
		if _, err := ToBoolE(value); err == nil {
			return boolType
		}
	}

	// Then check numeric types
	if intVal, err := ToIntE(value); err == nil {
		// Check if it's a string or another type that could be a float
		if floatVal, err := ToFloat64E(value); err == nil {
			// If the float value is equal to its integer part, it's an integer
			if float64(intVal) == floatVal {
				return int64Type
			} else {
				return float64Type
			}
		}
		return int64Type
	}

	// Check if it's an unsigned integer
	if _, err := ToUintE(value); err == nil {
		return uint64Type // Always return uint64 for unsigned integers
	}

	// Check if it's a float
	if _, err := ToFloat32E(value); err == nil {
		return float32Type
	}
	if _, err := ToFloat64E(value); err == nil {
		return float64Type
	}

	// Check if it's a string last
	if _, err := ToStringE(value); err == nil {
		return stringType
	}

	return InvalidType
}
