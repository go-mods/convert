package convert

import (
	"github.com/golang-module/carbon/v2"
	"reflect"
	"time"
)

// ToTime converts any type of value to time
func ToTime(value interface{}) (res time.Time, err error) {
	if reflect.TypeOf(value) == timeType {
		return value.(time.Time), nil
	} else {
		c := carbon.Parse(ToValidString(value))
		t := c.Carbon2Time()
		return t, c.Error
	}
}

// ToLayoutTime converts any type of value to time with a layout applied
func ToLayoutTime(layout string, value interface{}) (res time.Time, err error) {
	c := carbon.ParseByFormat(ToValidString(value), layout)
	if c.Error != nil {
		return time.Time{}, err
	}
	t := c.Carbon2Time()
	return t, c.Error
}

func ToTimeString(t time.Time, format ...string) (string, error) {
	c := carbon.Time2Carbon(t)
	if c.Error != nil {
		return "", c.Error
	}
	if len(format) > 0 {
		r := t.Format(format[len(format)-1])
		if r != format[len(format)-1] {
			return r, nil
		}
		return c.ToFormatString(format[len(format)-1]), nil
	}
	return c.String(), nil
}

// ToDuration converts any type of values to duration
func ToDuration(value interface{}) (res time.Duration, err error) {
	valueStr := ToValidString(value)
	res, err = time.ParseDuration(valueStr)
	return
}
