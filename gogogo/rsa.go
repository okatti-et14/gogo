package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func createRSA() {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}
	derRsaPrivateKey := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	f, err := os.Create("derFormatRsaPrivate.key")
	_, err = f.Write(derRsaPrivateKey)
	err = f.Close()
	f, err = os.Create("Private.key")
	err = pem.Encode(f, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: derRsaPrivateKey})
	err = f.Close()

	var rsaPublicKey crypto.PublicKey
	rsaPublicKey = rsaPrivateKey.Public()
	var derRsaPublicKey []byte
	if rsaPublicKeyPointer, ok := rsaPublicKey.(*rsa.PublicKey); ok {
		derRsaPublicKey, _ = x509.MarshalPKIXPublicKey(rsaPublicKeyPointer)
	}
	f, err = os.Create("derFormatRsaPublic.key")
	_, err = f.Write(derRsaPublicKey)
	err = f.Close()
	f, err = os.Create("Public.key")
	err = pem.Encode(f, &pem.Block{Type: "PUBLIC KEY", Bytes: derRsaPublicKey})
	err = f.Close()
}
