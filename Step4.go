package main

import "fmt"

type book = map[string]int
func main() {
  x := make(book,2)
  x["first"] = 777
  x["second"] = 666
  delete(x, "first")
  
  fmt.Println(x)
  fmt.Println(x["third"])

}