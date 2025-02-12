package main

import (
	"io"
	"strings"
)

func main() {
	str := "Hello, world\n"
	reader := strings.NewReader(str)

	io.ReadAll(reader)
}
