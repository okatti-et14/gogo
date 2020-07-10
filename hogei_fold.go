package main

import (
	"fmt"
	"strings"
)

func main() {
	prefix := "Basic "
	auth := "Basic eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFhIiwicGFzc3dvcmQiOiJhYSJ9.9mQjN4VUjvrTYOCm98DvXO1a_bPYAad2iMoNT9kS8E4"
	fmt.Println(len(prefix))
	fmt.Println(auth[:len(prefix)])
	fmt.Println(strings.EqualFold(auth[:len(prefix)], prefix))
	fmt.Println(!strings.EqualFold(auth[:len(prefix)], prefix))
	"jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1OTM2NTk4NTYsIm5iZiI6MTQ0NDQ3ODQwMCwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.XelBXbRNCUkItT8_7xpAip8fQ8NozQxpKAZQ125m06s",
	"api_key": "my_awesome_api_key",
	"oauth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1OTM2NTk4NTYsIm5iZiI6MTQ0NDQ3ODQwMCwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.XelBXbRNCUkItT8_7xpAip8fQ8NozQxpKAZQ125m06s"
}
