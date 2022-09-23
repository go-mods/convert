package convert

import "time"

// ToTime converts any type of values to time
func ToTime(value interface{}) (res time.Time, err error) {
	valueStr := ToValidString(value)
	res, err = time.Parse(time.RFC3339, valueStr)
	return
}

// ToLayoutTime converts any type of values to time
func ToLayoutTime(layout string, value interface{}) (res time.Time, err error) {
	valueStr := ToValidString(value)
	res, err = time.Parse(layout, valueStr)
	return
}

// ToDuration converts any type of values to duration
func ToDuration(value interface{}) (res time.Duration, err error) {
	valueStr := ToValidString(value)
	res, err = time.ParseDuration(valueStr)
	return
}
