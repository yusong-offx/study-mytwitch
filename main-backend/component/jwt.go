package component

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtkey = []byte("Hello-World-jwt-key")

func GernerateJWT(user_id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":    user_id,
		"expired": time.Now().Add(time.Minute * 15).Unix(),
	})
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtkey, nil
	})
	if err != nil {
		return nil, false
	}
	if c := token.Claims.(jwt.MapClaims); token.Valid {
		if v := c["expired"].(float64); int64(v) > time.Now().Unix() {
			return c, true
		}
		return nil, false
	}
	return nil, false
}
