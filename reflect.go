package convert

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"
)

type Caster func(value interface{}, asPtr bool) (reflect.Value, error)

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
	timePtrType  = reflect.TypeOf((*time.Time)(nil))
	durationType = reflect.TypeOf(time.Duration(0))
	nilType      = reflect.TypeOf(nil)
	InvalidType  = reflect.TypeOf(reflect.Value{})

	InvalidValue = reflect.Value{}
)

var casters = map[reflect.Type]Caster{
	reflect.TypeOf(""):               castString,
	reflect.TypeOf(false):            castBool,
	reflect.TypeOf(0):                castInt,
	reflect.TypeOf(int8(0)):          castInt8,
	reflect.TypeOf(int16(0)):         castInt16,
	reflect.TypeOf(int32(0)):         castInt32,
	reflect.TypeOf(int64(0)):         castInt64,
	reflect.TypeOf(uint(0)):          castUint,
	reflect.TypeOf(uint8(0)):         castUint8,
	reflect.TypeOf(uint16(0)):        castUint16,
	reflect.TypeOf(uint32(0)):        castUint32,
	reflect.TypeOf(uint64(0)):        castUint64,
	reflect.TypeOf(float32(0)):       castFloat32,
	reflect.TypeOf(float64(0)):       castFloat64,
	reflect.TypeOf(time.Now()):       castTime,
	reflect.TypeOf(time.Duration(0)): castTimeDuration,
}

func ToValue(value interface{}, to reflect.Type) (reflect.Value, error) {
	asPtr := to.Kind() == reflect.Pointer
	if asPtr {
		to = to.Elem()
	}

	if caster, ok := casters[to]; ok {
		return caster(value, asPtr)
	}

	return InvalidValue, errors.New("casting not supported")
}

func ToJsonValue(value interface{}, to reflect.Type) (reflect.Value, error) {
	asPtr := to.Kind() == reflect.Pointer
	if asPtr {
		to = to.Elem()
	}

	jsonString, err := ToString(value)
	if err != nil {
		return InvalidValue, err
	}

	jsonValuePtr := reflect.New(to)
	jsonValue := jsonValuePtr.Elem()

	err = json.Unmarshal([]byte(jsonString), jsonValue.Addr().Interface())
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return jsonValuePtr, nil
	}

	return jsonValue, nil
}

func castString(value interface{}, asPtr bool) (reflect.Value, error) {
	var s string
	var err error

	if reflect.TypeOf(value) == timeType {
		s = value.(time.Time).Format(time.RFC3339)
	} else {
		s, err = ToString(value)
		if err != nil {
			return InvalidValue, err
		}
	}

	if asPtr {
		return reflect.ValueOf(&s), nil
	}

	return reflect.ValueOf(s), nil
}

func castBool(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToBool(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castInt(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToInt(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castInt8(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToInt8(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castInt16(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToInt16(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castInt32(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToInt32(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castInt64(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToInt64(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castUint(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToUint(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castUint8(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToUint8(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castUint16(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToUint16(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castUint32(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToUint32(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castUint64(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToUint64(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castFloat32(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToFloat32(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castFloat64(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToFloat64(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castTime(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToTime(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func castTimeDuration(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToDuration(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
}

func GetConvertType(value interface{}) reflect.Type {
	if value == nil {
		return nil
	}

	// try to find the type
	if _, err := ToTime(value); err == nil {
		return timeType
	}
	if _, err := ToDuration(value); err == nil {
		return durationType
	}
	if _, err := ToInt(value); err == nil {
		if _, err := ToInt8(value); err == nil {
			if _, err := ToInt16(value); err == nil {
				if _, err := ToInt32(value); err == nil {
					if _, err := ToInt64(value); err == nil {
						return int64Type
					}
					return int32Type
				}
				return int16Type
			}
			return int8Type
		}
		return intType
	}
	if _, err := ToUint(value); err == nil {
		if _, err := ToUint8(value); err == nil {
			if _, err := ToUint16(value); err == nil {
				if _, err := ToUint32(value); err == nil {
					if _, err := ToUint64(value); err == nil {
						return uint64Type
					}
					return uint32Type
				}
				return uint16Type
			}
			return uint8Type
		}
		return uintType
	}
	if _, err := ToFloat32(value); err == nil {
		if _, err := ToFloat64(value); err == nil {
			return float64Type
		}
		return float32Type
	}
	if _, err := ToBool(value); err == nil {
		return boolType
	}
	if _, err := ToString(value); err == nil {
		return stringType
	}

	return InvalidType
}
