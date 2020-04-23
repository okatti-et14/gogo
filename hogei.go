package main

import (
	"fmt"
	"reflect"
)

//Takumi a of int , b of int
type Takumi struct {
	a int
	B int
}

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
	takumi := Takumi{1, 2}
	fmt.Println(reflect.TypeOf(takumi))
	fmt.Println(reflect.TypeOf(takumi.a))
}
