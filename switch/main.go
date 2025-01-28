package main

import (
  "fmt"
  "time"
  "math"
)

func main() {
  do(1)
  do(2)
  do(3)
  fmt.Println(isWeekend(time.Now()))
  switch x:= math.Sqrt(4); x{
    case 2:
      fmt.Println("resultado é 2")
    default:
      fmt.Println("algo deu errado")
  }
  fmt.Println(isNotWeekend(time.Now()))
  typeSwitch(1)
  typeSwitch("String Example")
  typeSwitch(1.0)
}

func typeSwitch(x any){
  switch x.(type) {
    case int:
      takeInt(x.(int))
      fmt.Println("int")
    case string:
      takeString(x.(string))
      fmt.Println("string")
    default:
      takeAny()
      fmt.Println("unknown")
  }
}

func takeString(x string){
  fmt.Println(x)
}

func takeInt(x int){
  fmt.Println(x)
}

func takeAny(){
  fmt.Println("any")
}

func do(x int){
  switch x {
    case 1:
      fmt.Println(1)
      fallthrough
    case 2:
      fmt.Println(2)
    default:
      fmt.Println("other")
  }
}

func isNotWeekend(x time.Time) bool {
  switch x.Weekday() {
    case time.Saturday, time.Sunday:
      return false
    default:
      return true
  }
}

func isWeekend(x time.Time) bool {
  switch {
  case x.Weekday() >0 && x.Weekday() < 6:
    return false
  default:
    return true
  }
}

