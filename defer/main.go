package main

import "fmt"

func main() {
  x := doDefer()
  fmt.Println(x)
  deferLifo()
  deferParam()
}

func doDefer() int {
  defer fmt.Println("World")
  fmt.Println("Hello")
  return 10
}

func deferLifo() {
  defer fmt.Println(3)
  defer fmt.Println(2)
  fmt.Println(1)
}

func deferParam() {
  x := 10
  defer func(y int) {
    fmt.Println(y)
  }(x)

  x = 50
  fmt.Println(x)
}
