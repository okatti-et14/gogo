package main

import (
	"fmt"
	"regexp"
)

func regxpExe() {
	//r, _ := regexp.Compile("{[a-z]*ch")
	r, _ := regexp.Compile(`.*ch`)
	fmt.Println(r.MatchString("{peach"))
}
