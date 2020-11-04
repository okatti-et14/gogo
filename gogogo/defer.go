package main

import "fmt"

func tmpdefer() {
	if true {
		fmt.Println("a")
		defer fmt.Println("b")
	}
	fmt.Println("c")
	defer fmt.Println("d")

}
