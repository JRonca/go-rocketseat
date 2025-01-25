package main

import (
	"fmt"
)

func main() {
	digaOi()
	fmt.Println(subtrair(1, 2))
	a, b := swap(1, 2)
	fmt.Println(a, b)
	res, rem := dividir(10, 3)
	fmt.Println(res, rem)
	f := somar(2)
	x := f(3)
	fmt.Println(x)
	fmt.Println(somarNums(1, 2, 3, 4, 5))
}

func digaOi() {
	fmt.Println("oi")
}

func subtrair(a, b int) int {
	return a - b
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

func somar(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func somarNums(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}
