package auth

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)


func Validate_token(token_str string) (bool) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err loading env", err)
	}

	key := os.Getenv("SECRET_KEY")

	if token_str == "" {
		return false
	}

	token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil || !token.Valid {
		log.Fatal("invalid token", err)
		return false
	}

	return true
}

