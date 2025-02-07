package main

import (
	"fmt"
	"math"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string {
	return s.msg
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"Cannot square root negative numbers"}
	}
	result := math.Sqrt(x)

	return result, nil
}

func printSquareRoot(x float64) {
	res, err := squareRoot(x)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func main() {
	x := -10.0
	y := 25.0
	printSquareRoot(x)
	printSquareRoot(y)
}
