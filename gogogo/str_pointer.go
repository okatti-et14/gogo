package main

import "fmt"

func checkStrpi() {
	ip := "127.0.0.1"
	ipp := &ip
	if ipp != (*string)(nil) {
		fmt.Println("true")
	}
}
