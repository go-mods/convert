package convert_test

import (
	"github.com/go-mods/convert"
	"testing"
)

func TestToJson(t *testing.T) {

	type Person struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		FullName  string `json:"full_name"`
	}
	var person = &Person{FirstName: "first", LastName: "last", FullName: "full"}
	var personStr = "{\"first_name\":\"first\",\"last_name\":\"last\",\"full_name\":\"full\"}"
	var personIndenStr = "{\n  \"first_name\": \"first\",\n  \"last_name\": \"last\",\n  \"full_name\": \"full\"\n}"

	personConvertStr, err := convert.ToJsonString(person)
	if err != nil {
		t.Errorf("ToJsonString() error = %v", err)
		return
	}

	if personStr != personConvertStr {
		t.Errorf("ToJson() got = %v, want %v", personConvertStr, personStr)
	}

	personConvertIndentStr, err := convert.ToJsonIndentString(person)
	if err != nil {
		t.Errorf("ToJsonIndentString() error = %v", err)
		return
	}

	if personIndenStr != personConvertIndentStr {
		t.Errorf("ToJson() got = %v, want %v", personConvertIndentStr, personIndenStr)
	}
}
