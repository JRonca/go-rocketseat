package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func main() {
	err := foo()
	if err != nil && errors.Is(err, ErrNotFound) {
		fmt.Println("Ok: Error is of type ErrNotFound")
		return
	}
	fmt.Println("No: Error is not of type ErrNotFound")
}

func foo() error { return ErrNotFound }
