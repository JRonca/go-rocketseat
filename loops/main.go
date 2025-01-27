package main

import "fmt"

func main() {
  var res int
  for i := 0; i < 10; i++ {
    res++
  }
  fmt.Println(res)

  // For range

  arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  for i, elem := range arr {
    fmt.Println(i, elem)
  }
  for i := range 10 {
    fmt.Println(i)
  }
}
