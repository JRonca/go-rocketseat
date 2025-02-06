package main

import "fmt"

func main() {
	const x = 10
	arrX := [x]string{4: "Hello, World!"}
	arr := [3]int{1, 2}
	arr1 := [10]int{5: 666, 7: 500}

	fmt.Println(arrX)
	fmt.Println(arr)
	fmt.Println(arr1)
}
