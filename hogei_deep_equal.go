package main

import (
	"fmt"
)

type tmp struct {
	Str  *string
	Int  *int
	hoge []*hoge
}

type hoge struct {
	Str *string
}

func main() {
	tmp := tmp{hoge: []*hoge{{}}}
	if tmp.Int == nil {
		fmt.Println("a")
	}
	if tmp.Str == nil {
		fmt.Println("b")
	}
	if tmp.hoge == nil {
		fmt.Println("c")
	}
	if tmp.hoge[0] == nil {
		fmt.Println("d")
	}
	if tmp.hoge[0].Str == nil {
		fmt.Println("e")
	}

}
