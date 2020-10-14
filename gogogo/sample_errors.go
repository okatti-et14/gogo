package main

import (
	"errors"
	"fmt"
	"runtime"
)

func buildErrors() {
	fmt.Println(Wrap(eeerror()))
}

func Wrap(err error) error {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	frame, _ := runtime.CallersFrames(pc).Next()
	s := fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
	return fmt.Errorf("%w %s", err, s)
}

func eeerror() error {
	err := errors.New("oirefjier")
	return Wrap(err)
}
