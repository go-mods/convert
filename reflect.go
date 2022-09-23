package convert

import (
	"errors"
	"reflect"
	"time"
)

type Caster func(value interface{}, asPtr bool) (reflect.Value, error)

var (
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

func castString(value interface{}, asPtr bool) (reflect.Value, error) {
	v, err := ToString(value)
	if err != nil {
		return InvalidValue, err
	}

	if asPtr {
		return reflect.ValueOf(&v), nil
	}

	return reflect.ValueOf(v), nil
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
