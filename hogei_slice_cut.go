package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a[0:1]
	c := a[0:6]
	d := a[0:0]
	fmt.Println(len(a))
	fmt.Println(a)
	fmt.Println(len(b))
	fmt.Println(b)
	fmt.Println(len(c))
	fmt.Println(c)
	fmt.Println(len(d))
	fmt.Println(d)
}
