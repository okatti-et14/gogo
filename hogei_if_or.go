package main

import "fmt"

func main() {

	var s *string

	if s == nil || *s == "" {
		fmt.Println("短絡評価だよ")
	}
}
