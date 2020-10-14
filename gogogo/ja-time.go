package main

import (
	"fmt"
	"time"
)

func outtime() {
	/* //utc, _ := time.LoadLocation("UTC")
	t := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	t = t.In(jst)
	tu := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)

	tj := time.Date(tu.Year(), tu.Month(), tu.Day(), 0, 0, 0, 0, jst)
	tjj := tj.In(jst)
	tj = tu.In(jst)
	fmt.Println(tjj.Hour())
	fmt.Println(tj.Hour()) */
	//t := time.Now()
	/* 	utc, _ := time.LoadLocation("UTC")
	   	jst, _ := time.LoadLocation("Asia/Tokyo")
	   	//tu := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, utc)
	   	fmt.Println(time.Now().In(jst))
	   	fmt.Println(time.Now().In(utc)) */

	/* 	before := time.Now()
	   	time.Sleep(2)
	   	after := time.Now()
	   	fmt.Println(before.After(after))
		   fmt.Println(after.After(before)) */
	utc, _ := time.LoadLocation("UTC")
	jst, _ := time.LoadLocation("Asia/Tokyo")
	a := time.Date(1995, 12, 12, 0, 0, 0, 0, utc)
	b := time.Date(1995, 12, 12, 0, 0, 0, 0, jst)
	fmt.Println(a)
	fmt.Println(a.In(jst))
	fmt.Println(b)
	fmt.Println(b.In(jst))
}
