package main

import "fmt"

func main() {
	foo("Teste")
	foo(123)
	foo([]int{1})
}

func foo[T any](arg T) {
	fmt.Println(arg)
}
