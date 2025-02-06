package pacote

import (
	"fmt"
	"packages/pacote/internal/foo"
)

var Bar string = "hello, Bar!"

func PrintMinha() {
	fmt.Println(foo.Minha)
}
