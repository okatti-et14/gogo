package main

import "fmt"

func zeroume() {
	s := fmt.Sprintf("%02d", 11)
	fmt.Println(s)
	s = fmt.Sprintf("%02d", 111)
	fmt.Println(s)
	s = fmt.Sprintf("%02d", 0)
	fmt.Println(s)
	s = fmt.Sprintf("%02d", 01)
	fmt.Println(s)
}
