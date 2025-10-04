package main

import (
  "fmt"
  "errors"
  )

func marsAge(age int) (int, error) {
  if age <= 0 {
    return age, errors.New("(AGE ERROR)")
  }
  return age*365/687, nil
}

func main() {
  var a int
  for {
    fmt.Scanln(&a)
    age, err := marsAge(a)
    if err != nil {
      panic("Ошибка ввода возраста")
    }
    a = age
    break
  }
  fmt.Println(a)
}