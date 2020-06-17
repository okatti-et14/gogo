package main

import "fmt"

func main() {
	collection := make([]*int, 1, 1)
	fmt.Println("1のとき↓")
	for index, val := range collection {
		fmt.Println(index)
		fmt.Println(val)
	}
	collection = nil
	fmt.Println("nilのとき↓")
	for index, val := range collection {
		fmt.Println(index)
		fmt.Println(val)
	}

	col2 := make([]*int, 0, len(collection))
	fmt.Println(len(collection))
	fmt.Println(col2)
}
