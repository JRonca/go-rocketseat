package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "Hello, world\n"
	reader := strings.NewReader(str)

	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Print(string(data))

	str2 := "abc"
	reader2 := strings.NewReader(str2)

	dataFull, errFull := io.ReadFull(reader2, make([]byte, 3))
	if errFull != nil {
		fmt.Println("Error reading: ", errFull)
		return
	}
	fmt.Print(dataFull)
}
