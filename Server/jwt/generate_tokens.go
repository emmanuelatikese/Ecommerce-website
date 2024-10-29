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
	accessClaims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 60 * 15).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}

	accToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	accessTokenStr, err := accToken.SignedString(accessTokenPrivate)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	refreshClaim := &jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}

	refToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaim)
	refreshTokenStr, err := refToken.SignedString(refreshTokenPrivate)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func GenerateAccessToken(id interface{}) (string, error) {
	accessClaims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 60 * 15).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}

	accToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	accessTokenStr, err := accToken.SignedString(accessTokenPrivate)
	if err != nil {
		log.Fatal(err)
		return "",  err
	}

	return accessTokenStr, nil
}