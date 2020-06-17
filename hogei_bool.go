package main

import (
	"fmt"
)

func main() {
	var i *int
	j := 1
	i = &j
	i = nil
	if i != nil && *i == 1 {
		fmt.Println(*i)
	} else {
		fmt.Println("nil")
	}
}
