package main

import "fmt"

func main() {
	cases := map[string]struct {
		id interface{}
	}{
		"normal": {
			id: 1,
		},
		"nil": {},
	}
	for _, c := range cases {

		PrintPointer(convertIntPointer(c.id))
	}

}

func PrintPointer(id *int) {
	if id == nil {
		return
	}
	fmt.Println(*id)
}

func convertIntPointer(i interface{}) *int {
	fmt.Println("i", i)
	fmt.Println(nil)
	num, ok := i.(int)
	if !ok {
		return nil
	}
	return &num
}
