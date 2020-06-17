package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User is
type User struct {
	ID   int
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
	fmt.Println(u.ID)
}
