package main

import (
	"fmt"
	"strings"
)

func splitExa() {
	errMessage := `value of userType must be one of "standard", "support" but got value "standard1234"`
	s := strings.TrimPrefix(errMessage, "value of")
	fmt.Println(s)
	params := strings.Split(s, " but got value ")

	transSlice := strings.Split(params[0], " must be one of ")
	fmt.Println("fodfijafiojeffoaflm", params[0])
	fmt.Println("fodfijafiojeffoaflm", transSlice[0])
}
