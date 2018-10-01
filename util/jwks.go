package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat/go-jwx/jwk"
)

func GetKey(region string, userPoolId string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolId)

		fmt.Printf("JWKS URL: %s", jwksURL)
		// TODO: cache response so we don't have to make a request every time
		// must be set manually in API Gateway Authorizers (checkbox)
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
}
