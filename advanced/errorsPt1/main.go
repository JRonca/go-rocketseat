package main

import (
	"errors"
	"fmt"
)

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func printDivision(a, b float64) {
	res, err := division(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func main() {
	a := 10.0
	b := 0.0
	c := 5.0
	printDivision(a, b)
	printDivision(a, c)
}
