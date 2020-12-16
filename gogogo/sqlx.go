package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Dummy struct {
	ID    int     `db:"id"`
	Dummy *string `db:"dummy"`
}

func Select() {
	sql := `select * from dummy where dummy is null;`
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		"postgres",
		"postgres",
		"localhost",
		5432,
		"postgres",
		"disable")
	conDB, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	dummy := []*Dummy{}
	conDB.Select(&dummy, sql)
	fmt.Println(len(dummy))
}

func SS() {
	sql := `select id from dummy;`
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		"postgres",
		"postgres",
		"localhost",
		5432,
		"postgres",
		"disable")
	conDB, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	result := []int{}

	conDB.Select(&result, sql)
	fmt.Println(len(result))
	fmt.Println(result)
	sort.Ints(result)
	fmt.Println(result)
	fmt.Println(result)
}
