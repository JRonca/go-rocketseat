package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]string{"a", "b", "c"}, "d"))
	fmt.Println(Contains([]string{"a", "b", "c"}, "b"))
	fmt.Println(slices.Contains([]int{1, 2, 3}, 2))
	fmt.Println(slices.Contains([]string{"a", "b", "c"}, "d"))
	fmt.Println(slices.Contains([]string{"a", "b", "c"}, "b"))
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
