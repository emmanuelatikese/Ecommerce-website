package jwt_util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)


func SetToken(priStr string, pubStr string) (*rsa.PrivateKey, *rsa.PublicKey, error){
	keyPrivate, err := os.ReadFile(priStr)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	block, _ := pem.Decode(keyPrivate)

	PriKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	keyPublic, err := os.ReadFile(pubStr)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	blockPub, _ := pem.Decode(keyPublic)

	PubKey, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	return PriKey.(*rsa.PrivateKey), PubKey.(*rsa.PublicKey), nil

}
