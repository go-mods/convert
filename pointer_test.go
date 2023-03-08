package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPtr_bool(t *testing.T) {
	var value bool = true
	var fallback bool = false

	result := ToPtr(value)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, FromPtrCb(result, fallback))
	assert.Exactly(t, fallback, FromPtrCb(nil, fallback))
}

func TestToPtr_string(t *testing.T) {
	var value string = "test"
	var fallback string = "fallback"

	result := ToPtr(value)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, FromPtrCb(result, fallback))
	assert.Exactly(t, fallback, FromPtrCb(nil, fallback))
}

func TestToPtr_int(t *testing.T) {
	var value int = 123
	var fallback int = 456

	result := ToPtr(value)

	assert.NotNil(t, result)
	assert.Exactly(t, &value, result)
	assert.Exactly(t, value, *result)

	assert.Exactly(t, value, FromPtrCb(result, fallback))
	assert.Exactly(t, fallback, FromPtrCb(nil, fallback))
}
