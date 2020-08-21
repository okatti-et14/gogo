package main

import (
	"fmt"
	"sort"
)

type Tmp struct {
	ID   *int
	Name *string
	Age  *int
}

func main() {
	i12 := 12
	i11 := 11
	saac := "aac"
	saab := "aab"
	saaa := "aaa"
	i110 := 110
	i100 := 100
	tmp := []*Tmp{
		{ID: &i12, Name: &saaa, Age: &i110},
		{ID: &i12, Name: &saac, Age: &i100},
		{ID: &i12, Name: &saab, Age: &i100},
		{ID: &i11, Name: &saaa, Age: &i100}}
	sort.SliceStable(tmp, func(i, j int) bool {
		return *tmp[i].Age < *tmp[j].Age
	})
	sort.SliceStable(tmp, func(i, j int) bool {
		return *tmp[i].Name < *tmp[j].Name
	})
	sort.SliceStable(tmp, func(i, j int) bool {
		return *tmp[i].ID < *tmp[j].ID
	})
	for _, b := range tmp {
		fmt.Println("id:", *b.ID, ",name:", *b.Name, ",age:", *b.Age)
	}
	fmt.Print(tmp)
}
