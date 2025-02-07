package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrQualquer))
	fmt.Println(errors.Is(err, ErrQualquer2))
}

var (
	ErrQualquer  = errors.New("error")
	ErrQualquer2 = errors.New("error 2")
)

func a() error { return ErrQualquer }
func b() error { return ErrQualquer2 }

func foo() error {
	var errorResult error
	if err := a(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}
	if err := b(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	return errorResult
}
