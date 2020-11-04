package main

import (
	"encoding/base64"
	"fmt"
)

func conBase64() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("ebisol")))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("expoiredpass")))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("ebica0000")))
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("expiredpass")))
}
