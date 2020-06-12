package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	a := [10000000]string{}
	for j := 0; j < 10000000; j++ {
		a[j] = "11"
	}
	b := [10000000]int{}
	c := make([]int, 0, 10000000)
	var err error
	start := time.Now()
	for i := range a {
		b[i], err = strconv.Atoi(a[i])
		if err != nil {
			fmt.Print(err)
		}
	}
	fmt.Println(time.Now().Sub(start))
	start = time.Now()
	d := 0
	for i := range a {
		d, err = strconv.Atoi(a[i])
		c = append(c, d)
	}
	fmt.Println(time.Now().Sub(start))
	fmt.Println(a[0], b[1])
	e := make([]int, 0, 2)
	f := make([]int, 0)
	g := []int{}
	ee := (*reflect.SliceHeader)(unsafe.Pointer(&e))
	ff := (*reflect.SliceHeader)(unsafe.Pointer(&f))
	gg := (*reflect.SliceHeader)(unsafe.Pointer(&g))
	fmt.Println("---------")
	fmt.Println("ee:", ee, ":", e)
	fmt.Println("ff:", ff, ":", f)
	fmt.Println("gg:", gg, ":", g)
	fmt.Println("---------")
	e = append(e, 1)
	f = append(f, 1)
	g = append(g, 1)
	fmt.Println("ee:", ee, ":", e)
	fmt.Println("ff:", ff, ":", f)
	fmt.Println("gg:", gg, ":", g)
	fmt.Println("---------")
	e = append(e, 1)
	f = append(f, 1)
	g = append(g, 1)
	fmt.Println("ee:", ee, ":", e)
	fmt.Println("ff:", ff, ":", f)
	fmt.Println("gg:", gg, ":", g)
	fmt.Println("---------")
}
