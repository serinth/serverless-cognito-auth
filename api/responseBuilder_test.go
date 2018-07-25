package api

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestBuilderCreatesCORSHeaders(t *testing.T) {
	r := NewResponseBuilder().
		Body("OK").
		Status(http.StatusOK).
		Build()

	assert.Equal(t, "*", r.Headers["Access-Control-Allow-Origin"], "Allow origin not set to *")
	assert.Equal(t, "true", r.Headers["Access-Control-Allow-Credentials"], "Allow credentials not set to [\"true\"]")
	assert.Equal(t, http.StatusOK, r.StatusCode, "HTTP status not set to 200")
	assert.Equal(t, false, r.IsBase64Encoded, "Default value of IsBase64Encoded is not false")
	assert.Equal(t, "OK", r.Body, "Body not set to \"OK\"")
}

func TestBuilderCustomHeader(t *testing.T) {
	r := NewResponseBuilder().
		AddHeader("MyCustomHeader", "MyCustomValue").
		Build()

	assert.Equal(t, "*", r.Headers["Access-Control-Allow-Origin"], "Allow origin not set to *")
	assert.Equal(t, "true", r.Headers["Access-Control-Allow-Credentials"], "Allow credentials not set to [\"true\"]")
	assert.Equal(t, "MyCustomValue", r.Headers["MyCustomHeader"], "[MyCustomHeader] did not have value: [MyCustomValue]")
}

func TestBuilderHTTPStatusCode(t *testing.T) {
	r := NewResponseBuilder().
		Status(http.StatusInternalServerError).
		Build()

	assert.Equal(t, http.StatusInternalServerError, r.StatusCode, "Custom status code did not return HTTP 500")
}

func TestIsBase64EncodedDefaultValue(t *testing.T) {
	r := NewResponseBuilder().
		IsBase64Encoded(true).
		Build()

	assert.Equal(t, true, r.IsBase64Encoded, "IsBase64Encoded not set to true")
}