package main

import "fmt"

func main() {
	var ms myStruct[string] = myStruct[string]{value: "Teste"}
	fmt.Println(ms)
}

type myStruct[T any] struct {
	value T
}
