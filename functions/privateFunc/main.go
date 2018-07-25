package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serinth/serverless-cognito-auth/api"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("events.APIGatewayProxyRequestContext.Identity: %#v \n", request.RequestContext)

	fmt.Printf("Headers: %#v \n", request.Headers)

	return api.NewResponseBuilder().
		Body("OK").
		Status(http.StatusOK).
		Build(), nil
}

func main() {
	lambda.Start(Handler)
}
