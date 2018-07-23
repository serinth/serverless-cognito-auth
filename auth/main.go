package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"github.com/lestrrat/go-jwx/jwk"
	"os"
	"strings"
)

const ENV_REGION = `REGION`
//TODO add configuration to get the jwksURL
func getKey(token *jwt.Token) (interface{}, error) {

	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json",os.Getenv(ENV_REGION), "<user pool id>")

	fmt.Printf("JWKS URL: %s", jwksURL)
	// TODO: cache response so we don't have to make a request every time
	// we want to verify a JWT
	set, err := jwk.FetchHTTP(jwksURL)
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	if key := set.LookupKeyID(keyID); len(key) == 1 {
		return key[0].Materialize()
	}

	return nil, errors.New("unable to find key")
}

// Help function to generate an IAM policy
func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}
	return authResponse
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	splitToken := strings.Split(event.AuthorizationToken, "Bearer ")
	tokenString := splitToken[1]

	token, err := jwt.Parse(tokenString, getKey)
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

	return generatePolicy(claims["sub"].(string), "Allow", event.MethodArn), nil
}

func main() {
	lambda.Start(handleRequest)
}