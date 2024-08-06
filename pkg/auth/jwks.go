package auth

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
	"log"
)

func FetchJwks(token *jwt.Token) (interface{}, error) {
	keySet, err := jwk.Fetch(context.Background(), "http://localhost:9000/v1/jwks")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	key, ok := keySet.Get(0)
	if !ok {
		return nil, errors.New("key not found")
	}
	var rawKey interface{}
	if err := key.Raw(&rawKey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}
	rsaKey, ok := rawKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Unable to parse to rsa key")
	}

	return rsaKey, nil
}
