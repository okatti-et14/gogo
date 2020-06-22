package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// User is
type User struct {
	ID       int        `json:"idid"`
	Name     *string    `json:"namename,omitempty"`
	BirthDay *time.Time `json:"birthday"`
}

func main() {
	hogeiUnmarshal()
	hogeiMarshal()
}

func hogeiUnmarshal() {
	src := `{"idid":3,"namename":"hogei","birthday":"2015-09-18T00:00:00+09:00"}`
	u := &User{}
	err := json.Unmarshal([]byte(src), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	fmt.Println(u.ID)
}

func hogeiMarshal() {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t := time.Date(2015, 9, 18, 0, 0, 0, 0, loc)
	u := User{
		BirthDay: &t,
		ID:       1,
	}

	j, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
