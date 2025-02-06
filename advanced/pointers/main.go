package main

import "fmt"

func main() {
	x := 10
	p := &x
	fmt.Println(p, *p)
	take(&x)
	fmt.Println(x)
}

func take(x *int) {
	*x = 100
}
