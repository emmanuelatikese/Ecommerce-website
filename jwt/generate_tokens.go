package jwt_util

import (
	"crypto/rsa"
	"log"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/subosito/gotenv"
)

var (
	access_token_private *rsa.PrivateKey
	refresh_token_private *rsa.PrivateKey

	Access_token_public *rsa.PublicKey
	Refresh_token_public *rsa.PublicKey
)

func init() {
	gotenv.Load()
	access_pri_dir, access_pub_dir := os.Getenv("ACCESS_PRIVATE_TOKEN_DIR"), os.Getenv("ACCESS_PUBLIC_TOKEN_DIR")
	refresh_pri_dir, refresh_pub_dir := os.Getenv("REFRESH_PRIVATE_TOKEN_DIR"), os.Getenv("REFRESH_PUBLIC_TOKEN_DIR")
	var err error
	access_token_private, Access_token_public, err = SetToken(access_pri_dir, access_pub_dir)
	if err != nil{
		log.Fatal(err)
		return
	}

	refresh_token_private, Refresh_token_public, err = SetToken(refresh_pri_dir, refresh_pub_dir)
	if err != nil{
		log.Fatal(err)
		return
	}

}

func GenerateToken(id interface{}) (string, string, error) {
	access_claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 60 * 15).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}

	acc_token := jwt.NewWithClaims(jwt.SigningMethodRS256, access_claims)
	access_token_str, err := acc_token.SignedString(access_token_private)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	refresh_claim := &jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}

	ref_token := jwt.NewWithClaims(jwt.SigningMethodRS256, refresh_claim)
	refresh_token_str, err := ref_token.SignedString(refresh_token_private)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	return access_token_str, refresh_token_str, nil
}

func GenerateAccessToken(id interface{}) (string, error) {
	access_claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 60 * 15).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}

	acc_token := jwt.NewWithClaims(jwt.SigningMethodRS256, access_claims)
	access_token_str, err := acc_token.SignedString(access_token_private)
	if err != nil {
		log.Fatal(err)
		return "",  err
	}

	return access_token_str, nil
}