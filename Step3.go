package main

import "fmt"


func main() {
  x := [4]int{12,6,9,3}
  y := nil
  fmt.Println(x)
  reverse(&x)
  fmt.Println(x)
  g := &y
  fmt.Println(*g)
}

func reverse(xMap *[4]int) {
  for key, value := range *xMap {
    xMap[len(xMap)-1-key] = value
  }
  // fmt.Println(xMap)
}