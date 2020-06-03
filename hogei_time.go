package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	time.Sleep(time.Second * 1)
	b := time.Now()
	c := b.Sub(a).Milliseconds()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(fmt.Sprintf("%vミリ秒", c))
}
