package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func devstrcase() {
	manyAAA := bytes.NewBufferString("AAA")
	for i := 0; i < 100000; i++ {
		manyAAA.WriteString("AAAAAA")
	}
	manyAAA2 := bytes.NewBufferString("AAA")
	for i := 0; i < 1000; i++ {
		manyAAA2.WriteString("AAAAAA")
	}

	now := time.Now()
	for i := 0; i < 10000; i++ {
		strings.Contains(manyAAA2.String(), manyAAA.String())
	}
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	for i := 0; i < 10000; i++ {
		strings.HasSuffix(manyAAA2.String(), manyAAA.String())
	}
	fmt.Println(time.Now().Sub(now))

}
