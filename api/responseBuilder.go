package api

import (
	"github.com/aws/aws-lambda-go/events"
)

type responseBuilder struct {
	response events.APIGatewayProxyResponse
}

func NewResponseBuilder() *responseBuilder {
	return &responseBuilder{}
}

func (rb *responseBuilder) Build() events.APIGatewayProxyResponse {
	rb.AddHeader("Access-Control-Allow-Origin", "*")
	rb.AddHeader("Access-Control-Allow-Credentials", "true")
	return rb.response
}

func (rb *responseBuilder) Body(body string) *responseBuilder {
	rb.response.Body = body
	return rb
}

func (rb *responseBuilder) AddHeader(key string, val string) *responseBuilder {
	if rb.response.Headers == nil {
		rb.response.Headers = map[string]string {key: val}
	} else {
		rb.response.Headers[key] = val
	}
	return rb
}

func (rb *responseBuilder) Status(code int) *responseBuilder {
	rb.response.StatusCode = code
	return rb
}

func (rb *responseBuilder) IsBase64Encoded(isBase64Encoded bool) *responseBuilder {
	rb.response.IsBase64Encoded = isBase64Encoded
	return rb
}