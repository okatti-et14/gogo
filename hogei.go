package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 5
	j := &i
	fmt.Println(i)
	fmt.Println(*j)
	fmt.Println(j)
	fmt.Println(&i)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(&i))
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(&j))
	fmt.Println(reflect.TypeOf(*j))
}
