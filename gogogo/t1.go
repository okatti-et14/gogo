package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func t1() {
	z := make([]int, 0, 10)
	aa := make(chan int, 10)
	eg, ctx := errgroup.WithContext(context.Background())
	ii := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range ii {
		i := i
		eg.Go(func() error {
			select {
			case <-ctx.Done():

				fmt.Println(i)
				return nil
			default:
				fmt.Println(i)

				if i == 2 {
					return errors.New("Error occurred")
				}
				aa <- i
				return nil
			}
		})
		select {
		case a := <-aa:
			z = append(z, a)
		case <-ctx.Done():
			break
		}
	}
	fmt.Println(z)

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Complete!")
}
