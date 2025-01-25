package main

import (
	"fmt"
	"strconv"
)

func main() {
	nome := "José"
	sobrenome := "Renato"

	fmt.Println(nome, sobrenome)

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(u)

	var x int = 10084
	s := string(rune(x))
	fmt.Println(s)

	newString := strconv.FormatInt(int64(x), 10)
	fmt.Println(newString)

	// constantes

	const xConst int = 10

	// ou

	const yConst = 20

	fmt.Println(xConst)
	fmt.Println(yConst)

	takeInt32(yConst)
	takeFloat64(yConst)
	takeString("Hello")
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeString(s string) {
	fmt.Println(s)
}

func takeFloat64(x float64) {
	fmt.Println(x)
}
