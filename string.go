package convert

import (
	"fmt"
	"reflect"
)

// ToString converts any type of value to string
func ToString(value interface{}) (string, error) {
	return ToValidString(value), nil
}

// ToValidString converts any type of value to string
func ToValidString(value interface{}) string {
	switch v := value.(type) {
	case fmt.Stringer:
		return v.String()
	case string:
		return v
	default:
		return fmt.Sprintf("%v", value)
	}
}

// ToStringArray converts any type of value to array of string
func ToStringArray(value interface{}) ([]string, error) {
	return ToValidStringArray(value), nil
}

// ToValidStringArray converts any type of value to array of string
func ToValidStringArray(array interface{}) (resArray []string) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(array)
			resArray = make([]string, v.Len())
			for index := range resArray {
				resArray[index] = ToValidString(v.Index(index).Interface())
			}
			return
		}
	}
	return []string{ToValidString(array)}
}
