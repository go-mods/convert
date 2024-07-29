package convert

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CustomType struct {
	Field string
}

func (c CustomType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"custom_field": c.Field})
}

func TestToJson(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	expected := `{"key":"value"}`
	result := ToJson(value)
	assert.JSONEq(t, expected, string(result))

	// Test with a custom converter
	customConverter := func(value interface{}) *[]byte {
		if v, ok := value.(CustomType); ok {
			b, err := json.Marshal(v)
			if err == nil {
				return &b
			}
		}
		return nil
	}
	customValue := CustomType{Field: "test"}
	expectedCustom := `{"custom_field":"test"}`
	resultCustom := ToJson(customValue, customConverter)
	assert.JSONEq(t, expectedCustom, string(resultCustom))

	// Test with nil value
	resultNil := ToJson(nil)
	assert.Nil(t, resultNil)
}

func TestToJsonOrDefault(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	defaultValue := []byte(`{}`)
	expected := `{"key":"value"}`
	result := ToJsonOrDefault(value, defaultValue)
	assert.JSONEq(t, expected, string(result))

	// Test with nil value
	resultDefault := ToJsonOrDefault(nil, defaultValue)
	assert.JSONEq(t, string(defaultValue), string(resultDefault))
}

func TestToJsonE(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	expected := `{"key":"value"}`
	result, err := ToJsonE(value)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(result))

	// Test with an invalid value
	invalidValue := make(chan int)
	_, err = ToJsonE(invalidValue)
	assert.Error(t, err)

	// Test with nil value
	resultNil, err := ToJsonE(nil)
	assert.Error(t, err)
	assert.Nil(t, resultNil)
}

func TestToJsonString(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	expected := `{"key":"value"}`
	result := ToJsonString(value)
	assert.JSONEq(t, expected, result)

	// Test with nil value
	resultNil := ToJsonString(nil)
	assert.Equal(t, "", resultNil)
}

func TestToJsonStringOrDefault(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	defaultValue := "{}"
	expected := `{"key":"value"}`
	result := ToJsonStringOrDefault(value, defaultValue)
	assert.JSONEq(t, expected, result)

	// Test with nil value
	resultDefault := ToJsonStringOrDefault(nil, defaultValue)
	assert.JSONEq(t, defaultValue, resultDefault)
}

func TestToJsonStringE(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	expected := `{"key":"value"}`
	result, err := ToJsonStringE(value)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, result)

	// Test with an invalid value
	invalidValue := make(chan int)
	_, err = ToJsonStringE(invalidValue)
	assert.Error(t, err)

	// Test with nil value
	resultNil, err := ToJsonStringE(nil)
	assert.Error(t, err)
	assert.Equal(t, "", resultNil)
}

func TestToJsonIndent(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	expected := `{
  "key": "value"
}`
	result := ToJsonIndent(value)
	assert.JSONEq(t, expected, string(result))

	// Test with nil value
	resultNil := ToJsonIndent(nil)
	assert.Nil(t, resultNil)
}

func TestToJsonIndentOrDefault(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	defaultValue := []byte(`{}`)
	expected := `{
  "key": "value"
}`
	result := ToJsonIndentOrDefault(value, defaultValue)
	assert.JSONEq(t, expected, string(result))

	// Test with nil value
	resultDefault := ToJsonIndentOrDefault(nil, defaultValue)
	assert.JSONEq(t, string(defaultValue), string(resultDefault))
}

func TestToJsonIndentE(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	expected := `{
  "key": "value"
}`
	result, err := ToJsonIndentE(value)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, string(result))

	// Test with an invalid value
	invalidValue := make(chan int)
	_, err = ToJsonIndentE(invalidValue)
	assert.Error(t, err)

	// Test with nil value
	resultNil, err := ToJsonIndentE(nil)
	assert.Error(t, err)
	assert.Nil(t, resultNil)
}

func TestToJsonIndentString(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	expected := `{
  "key": "value"
}`
	result := ToJsonIndentString(value)
	assert.JSONEq(t, expected, result)

	// Test with nil value
	resultNil := ToJsonIndentString(nil)
	assert.Equal(t, "", resultNil)
}

func TestToJsonIndentStringOrDefault(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	defaultValue := "{}"
	expected := `{
  "key": "value"
}`
	result := ToJsonIndentStringOrDefault(value, defaultValue)
	assert.JSONEq(t, expected, result)

	// Test with nil value
	resultDefault := ToJsonIndentStringOrDefault(nil, defaultValue)
	assert.JSONEq(t, defaultValue, resultDefault)
}

func TestToJsonIndentStringE(t *testing.T) {
	// Test with a simple map
	value := map[string]interface{}{"key": "value"}
	expected := `{
  "key": "value"
}`
	result, err := ToJsonIndentStringE(value)
	assert.NoError(t, err)
	assert.JSONEq(t, expected, result)

	// Test with an invalid value
	invalidValue := make(chan int)
	_, err = ToJsonIndentStringE(invalidValue)
	assert.Error(t, err)

	// Test with nil value
	resultNil, err := ToJsonIndentStringE(nil)
	assert.Error(t, err)
	assert.Equal(t, "", resultNil)
}
