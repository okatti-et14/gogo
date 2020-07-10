package main

import "fmt"

func main() {
	var a = [2]string{"hello", "takukmi"}
	b := [2]string{"hello", "takukmi"}
	fmt.Println(a[0], b[1])
	var c *int
	fmt.Print(c)
	gg := []int{1, 2, 3}
	fmt.Println(gg)
	base := []int{}
	i := []int{1, 2, 3}
	j := []int{}
	base = append(base, 1)
	base = append(base, i...)
	fmt.Println("base:", base)
	base = append(base, j...)
	fmt.Println("base:", base)
	fmt.Println("j", j)
}
