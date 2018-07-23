package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {


	fmt.Printf("events.APIGatewayProxyRequestContext.Identity: %#v \n", request.RequestContext)

	fmt.Printf("Headers: %#v \n", request.Headers)


	return events.APIGatewayProxyResponse{
		Body:       "authFunc called after authenticated",
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
