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
	accessTokenPrivate *rsa.PrivateKey
	refreshTokenPrivate *rsa.PrivateKey

	AccessTokenPublic *rsa.PublicKey
	RefreshTokenPublic *rsa.PublicKey
)

func init() {
	gotenv.Load()
	accessPriDir, accessPubDir := os.Getenv("ACCESS_PRIVATE_TOKEN_DIR"), os.Getenv("ACCESS_PUBLIC_TOKEN_DIR")
	refreshPriDir, refreshPubDir := os.Getenv("REFRESH_PRIVATE_TOKEN_DIR"), os.Getenv("REFRESH_PUBLIC_TOKEN_DIR")
	var err error
	accessTokenPrivate, AccessTokenPublic, err = SetToken(accessPriDir, accessPubDir)
	if err != nil{
		log.Fatal(err)
		return
	}
	refreshTokenPrivate, RefreshTokenPublic, err = SetToken(refreshPriDir, refreshPubDir)
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
	access_token_str, err := acc_token.SignedString(accessTokenPrivate)
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
	refresh_token_str, err := ref_token.SignedString(refreshTokenPrivate)
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
	access_token_str, err := acc_token.SignedString(accessTokenPrivate)
	if err != nil {
		log.Fatal(err)
		return "",  err
	}

	return access_token_str, nil
}