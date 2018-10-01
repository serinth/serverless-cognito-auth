package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/serinth/serverless-cognito-auth/util"
	"os"
	"strings"
)

var REGION string
var USER_POOL_ID string
var APP_CLIENT_ID string

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	splitToken := strings.Split(event.AuthorizationToken, "Bearer ")
	tokenString := splitToken[1]

	token, err := jwt.Parse(tokenString, util.GetKey(REGION, USER_POOL_ID))
	if err != nil {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New(fmt.Sprintf("Error: Invalid token with error: %v", err))
	}

	if !token.Valid {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized") // Return a 401 Unauthorized response
	}

	claims := token.Claims.(jwt.MapClaims)
	for key, value := range claims {
		fmt.Printf("%s\t%v\n", key, value)
	}

	badClaims := claims.Valid()

	if badClaims != nil || !claims.VerifyAudience(APP_CLIENT_ID, true) {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Invalid Token.")
	}

	return util.GenerateLambdaInvokePolicy(claims["sub"].(string), "Allow", event.MethodArn), nil
}

func init() {
	REGION = os.Getenv("REGION")
	USER_POOL_ID = os.Getenv("USER_POOL_ID")
	APP_CLIENT_ID = os.Getenv("APP_CLIENT_ID")
}

func main() {
	lambda.Start(handleRequest)
}
