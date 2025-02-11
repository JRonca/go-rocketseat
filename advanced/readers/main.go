package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "Hello, world"
	reader := strings.NewReader(str)

	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("Error at read buffer")
			return
		}
		fmt.Println(n)
		fmt.Println(buffer[:n])
		fmt.Println(string(buffer[:n]))
	}
}
