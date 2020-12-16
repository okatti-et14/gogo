package main

import "fmt"

func goroutineSample() {
	numchan := make(chan int, 1)
	go goAA(numchan)

	aa := <-numchan
	fmt.Println("before")
	fmt.Println(aa)
	fmt.Println("after")

}

func goAA(numchan chan int) {
	for i := 0; i < 10000; i++ {
		if i%333 == 0 {
			fmt.Println("now:", i)
		}
	}
	numchan <- 999
}
