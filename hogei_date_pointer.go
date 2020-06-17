package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local)
	pDate := &date
	//pDate = nil
	var pStrDate *string
	var strDate string
	if pDate != nil {
		strDate = pDate.Format("2006/01/02")
		pStrDate = &strDate
	}
	fmt.Println(pStrDate)
}
