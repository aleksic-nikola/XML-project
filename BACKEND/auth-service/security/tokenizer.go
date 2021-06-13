package security

import (
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)


func GetToken(username string, role string) (string, error) {
	godotenv.Load()
	signingKey := []byte(os.Getenv("ACCESS_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role": role,
		"expires" : time.Now().Add(time.Minute * 30).Unix(),
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func GerifyToken(tokenString string) (jwt.Claims, error) {
	godotenv.Load()
	signingKey := []byte(os.Getenv("ACCESS_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}