package jwt_util

import (
	"errors"
	"log"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
)

func Verify(w http.ResponseWriter, r *http.Request, value string) (string, error) {
	tk, err := jwt.Parse(value, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodRSA); !ok {
			err := errors.New("invalid token")
			return "", err
		}
		return Refresh_token_public, nil
	})
	if err != nil {
		log.Print(err)
		return "", err
	}
	if tk.Valid {
		if jwtMap, ok := tk.Claims.(jwt.MapClaims); ok {
			value := jwtMap["sub"]
			str_val, ok := value.(string)
			if !ok{
				log.Print(err)
				return "", err
			}
			return str_val, nil
		}
	}
	return "", err
}
