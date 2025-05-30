package auth

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func ExtractToken(authHeader string) string {
	const prefix = "Bearer "
	if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
		return authHeader[len(prefix):]
	}
	return authHeader
}

func ValidateToken(token_str string) bool {
	encodedKey := os.Getenv("SECRET_KEY")

	if token_str == "" {
		return false
	}

	raw_token := ExtractToken(token_str)

	key, err := base64.RawURLEncoding.DecodeString(encodedKey)
	if err != nil {
		log.Fatal("error decoding secret key:", err)
		return false
	}

	token, err := jwt.Parse(raw_token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil || !token.Valid {
		log.Println("invalid token:", err)
		return false
	}

	return true
}

