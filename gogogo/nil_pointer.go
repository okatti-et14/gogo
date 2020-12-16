package main

import "fmt"

type nilnil1 struct {
	s1 []*string1
}
type nilnil2 struct {
	s2 []*string2
}

type string1 struct {
	a string
	b *string
}
type string2 struct {
	a string
	b *string
}

func nilPointter() {
	/*aa := &nilnil1{
		s1: []*string1{
			{a: "a"},
		},
	}
	bb := &nilnil2{
		s2: []*string2{
			{a: "a"},
		},
	}

	fmt.Println("a:", aa.s1[0].a == bb.s2[0].a)
	fmt.Println("b:", aa.s1[0].b == bb.s2[0].b)
	fmt.Println(aa.s1[0].b)
	fmt.Println(bb.s2[0].b)*/
	var aa *string
	aaaaa(*aa)

}

func aaaaa(str string) {
	fmt.Println(str)
	fmt.Println(&str)
}
