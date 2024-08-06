package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type Role string

type JwtCustomClaims struct {
	Username string `json:"username"`
	Role     Role   `json:"role"`
	jwt.RegisteredClaims
}

func ExtractClaims(tokenStr string) (*JwtCustomClaims, bool) {
	publicKey, err := LoadRSAPublicKeyFromDisk(publicKeyPath)
	if err != nil {
		log.Fatalf("Error loading RSA public key: %v", err)
	}

	customClaims := &JwtCustomClaims{} // Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenStr, customClaims, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	tokenClaims, ok := token.Claims.(*JwtCustomClaims)
	return tokenClaims, ok
}

// Replace with the path to your RSA public key
const publicKeyPath = "pkg/auth/public.pem"
const privateKeyPath = "pkg/auth/private.pem"

func GenerateJwt(claims JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	key, err := LoadRSAPrivateKeyFromDisk(privateKeyPath)
	if err != nil {
		return "", err
	}

	// Generate encoded token and send it as response.
	t, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return t, nil

}
