package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	if err != nil && errors.Is(err, ErrQualquer) {
		fmt.Println(err)
		return
	}
}

func foo() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	return nil
}

var ErrQualquer = errors.New("error")

func bar() error {
	return ErrQualquer
}
