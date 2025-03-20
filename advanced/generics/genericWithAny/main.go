package main

import (
	"fmt"
)

func main() {
	foo("Teste Any")
	foo(123)
	foo([]int{1})
	fooComparable("Teste Comparable")
}

func foo[T any](arg T) {
	fmt.Println(arg)
}

func fooComparable[T comparable](arg T) {
	fmt.Println(arg)
}
