package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	str := "Hello, world\n"
	reader := strings.NewReader(str)
	writer := MyWriter{}

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

		_, _ = writer.Write(buffer[:n])
	}
}

type MyWriter struct{}

func (MyWriter) Write(b []byte) (int, error) {
	fmt.Print(string(b))
	return len(b), nil
}
