package main

import (
  "fmt"
  )

// type Car struct {
//   color, brand string
//   year int
// } 

func main() {
  x:=98
  if fmt.Sprintf("%T", x) == "int" {
    fmt.Println("x is int")
  }
}