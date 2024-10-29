package jwt_util

import (
	"errors"
	"log"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
)

func Verify(w http.ResponseWriter, r *http.Request, value string, publicToken interface{}) (string, error) {
	tk, err := jwt.Parse(value, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodRSA); !ok {
			err := errors.New("invalid token")
			return "", err
		}
		return publicToken, nil
	})
	if err != nil {
		log.Print(err)
		return "", err
	}
	if tk.Valid {
		if jwtMap, ok := tk.Claims.(jwt.MapClaims); ok {
			value := jwtMap["sub"]
			strVal, ok := value.(string)
			if !ok{
				log.Print(err)
				return "", err
			}
			return strVal, nil
		}
	}
	return "", err
}
