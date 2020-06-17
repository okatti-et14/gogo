package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id   *int
	Name *string
}

func main() {
	src := `{"Id":3,"Name":"hogei"}`
	u := &User{}
	err := json.Unmarshal([]byte(src), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	fmt.Println(u.Id)
	fmt.Println(*u.Id)
}
