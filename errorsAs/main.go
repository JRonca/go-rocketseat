package main

import (
	"errors"
	"fmt"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string {
	return s.msg
}

func main() {
	err := foo()
	var sqrtError SqrtError
	if err != nil && errors.As(err, &sqrtError) {
		fmt.Println(sqrtError.msg)
		return
	}
	fmt.Println("No: Error is not of type SqrtError")
}

func foo() error { return SqrtError{msg: "Test"} }
