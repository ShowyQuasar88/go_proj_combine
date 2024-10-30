package jwt

import "github.com/golang-jwt/jwt/v5"

var key = "asf1as165as891"

type Data struct {
	Name   string
	Age    int
	Gender int
	jwt.RegisteredClaims
}

func Sign(data jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	sign, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return sign, nil
}

func Verify(sign string, data jwt.Claims) error {
	_, err := jwt.ParseWithClaims(sign, data, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return err
}
