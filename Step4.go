package main

import "fmt"

func main() {
  x := 2
  double(&x)
  fmt.Println(x)
  double(&x)
  fmt.Println(x)
}

func double(x *int) {
  *x *= 2
}