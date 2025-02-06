package main

import "fmt"

func main() {
	// slice about array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	arr[1] = 20
	slice[1] = 25
	fmt.Println(arr)
	fmt.Println(slice)
	// slice literal, length, capacity
	sliceLiteral := []int{1, 2, 3}
	slice2 := sliceLiteral[:0]
	fmt.Println(sliceLiteral, len(sliceLiteral), cap(sliceLiteral))
	fmt.Println(slice2, len(slice2), cap(slice2))
	// slice empty and slice nil
	sliceEmpty := []int{}
	fmt.Println(sliceEmpty == nil)
	var sliceNil []int
	fmt.Println(sliceNil == nil)
	// slice append
	sliceAppend := []int{}
	for i := 0; i < 10; i++ {
		fmt.Println(len(sliceAppend), cap(sliceAppend))
		sliceAppend = append(sliceAppend, i+1)
	}
	fmt.Println(sliceAppend)
	// sice make
	sliceMake := make([]int, 0, 10)
	fmt.Println(sliceMake)
	for j := 0; j < 10; j++ {
		fmt.Println(len(sliceMake), cap(sliceMake))
		sliceMake = append(sliceMake, j+1)
	}
	fmt.Println(sliceMake)
	// Bound check
	sliceBound := []int{1, 2, 3, 4, 5}
	foo(sliceBound)
}

func foo(slice []int) {
	_ = slice[3]
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
}
