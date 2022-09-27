package convert

import "encoding/json"

func ToJson(i interface{}) (d []byte, err error) {
	d, err = json.Marshal(i)
	return
}

func ToJsonOrPanic(i interface{}) []byte {
	if res, err := ToJson(i); err == nil {
		return res
	} else {
		panic(err)
	}
}

func ToJsonString(i interface{}) (string, error) {
	var d []byte
	var err error
	d, err = ToJson(i)
	return string(d), err
}

func ToJsonStringOrPanic(i interface{}) string {
	if res, err := ToJsonString(i); err == nil {
		return res
	} else {
		panic(err)
	}
}

func ToJsonIndent(i interface{}) (d []byte, err error) {
	d, err = json.MarshalIndent(i, "", "  ")
	return
}

func ToJsonIndentOrPanic(i interface{}) string {
	if res, err := ToJsonIndent(i); err == nil {
		return string(res)
	} else {
		panic(err)
	}
}

func ToJsonIndentString(i interface{}) (string, error) {
	var d []byte
	var err error
	d, err = ToJsonIndent(i)
	return string(d), err
}

func ToJsonIndentStringOrPanic(i interface{}) string {
	if res, err := ToJsonIndentString(i); err == nil {
		return string(res)
	} else {
		panic(err)
	}
}
