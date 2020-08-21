package main

import (
	"fmt"
	"sort"
)

type Tmp struct {
	ID   int
	Name string
	Age  int
}

func main() {
	tmp := []Tmp{
		{ID: 12, Name: "aac", Age: 110},
		{ID: 12, Name: "aac", Age: 100},
		{ID: 12, Name: "aab", Age: 100},
		{ID: 11, Name: "aaa", Age: 100}}
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i].Age < tmp[j].Age
	})
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i].Name < tmp[j].Name
	})
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i].ID < tmp[j].ID
	})
	fmt.Print(tmp)
}
