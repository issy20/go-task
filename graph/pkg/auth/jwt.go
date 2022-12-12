package auth

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/issy20/go-task/graph/utils"
)

var (
	secret_key = []byte("secret")
)

func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = utils.CurrentTime().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(secret_key)
	if err != nil {
		log.Fatal("Error in Generating key : ", err)
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id := claims["user_id"].(string)
		return user_id, nil
	} else {
		return "", err
	}
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
