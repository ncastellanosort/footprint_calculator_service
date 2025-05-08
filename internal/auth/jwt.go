package auth

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func Extract_token(authHeader string) string {
    const prefix = "Bearer "
    if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
        return authHeader[len(prefix):]
    }
    return authHeader 
}


func Validate_token(token_str string) bool {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("err loading env", err)
	}

	key := os.Getenv("SECRET_KEY")

	if token_str == "" {
		return false
	}

	raw_token := Extract_token(token_str)

	token, err := jwt.Parse(raw_token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil || !token.Valid {
		log.Fatal("invalid token", err)
		return false
	}

	return true
}
