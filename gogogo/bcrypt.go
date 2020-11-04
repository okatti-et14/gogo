package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func testbcrypt() {
	a, _ := bcrypt.GenerateFromPassword([]byte("ebisol"+"1@e0140000001"), 10)
	fmt.Println(string(a))
	a, _ = bcrypt.GenerateFromPassword([]byte("expiredpass"+"2@e0140000001"), 10)
	fmt.Println(string(a))
	a, _ = bcrypt.GenerateFromPassword([]byte("expiredpass"+"3@e0140000001"), 10)
	fmt.Println(string(a))
	a, _ = bcrypt.GenerateFromPassword([]byte("expiredpass"+"4@e0140000001"), 10)
	fmt.Println(string(a))
	a, _ = bcrypt.GenerateFromPassword([]byte("expiredpass"+"5@e0140000001"), 10)
	fmt.Println(string(a))
	a, _ = bcrypt.GenerateFromPassword([]byte("ebisol"+"0@e0140000001"), 10)
	fmt.Println(string(a))
	a, _ = bcrypt.GenerateFromPassword([]byte("ebisol"+"1@e0140000002"), 10)
	fmt.Println(string(a))
	a, _ = bcrypt.GenerateFromPassword([]byte("ebisol"+"1@e0140000003"), 10)
	fmt.Println(string(a))

}
