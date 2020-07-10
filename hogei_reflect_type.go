package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	typeInfo(i)
	s := "aaa"
	typeInfo(s)
	var inter interface{}
	typeInfo(inter)

}

func typeInfo(hoge interface{}) {
	fmt.Println(reflect.TypeOf(hoge))
}
