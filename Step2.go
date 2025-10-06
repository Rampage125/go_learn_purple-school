package main

import (
  "fmt"
  )

type bookMap = map[string]string

func main() {
  bookmarks := bookMap{}
  var variant int
Menu:
  for {
    variant = choice()
    switch variant {
      case 1: fmt.Println(bookmarks)
      case 2: addBookmark(bookmarks)
      case 3: deleteBookmark(bookmarks)
      case 4: break Menu
    }
  }
}


func choice() int {
  var x int
  fmt.Println(`— 1. Посмотреть закладки
— 2. Добавить закладку
— 3. Удалить закладку
— 4. Выход`)
  fmt.Scan(&x)
  return x
}

func addBookmark(bookmarks bookMap) {
  var key, value string
  fmt.Println("Введите ключ: ")
  fmt.Scan(&key)
  fmt.Println("Введите значение: ")
  fmt.Scan(&value)
  bookmarks[key] = value
}

func deleteBookmark(bookmarks bookMap) {
  var key string
  fmt.Println("Введите ключ: ")
  fmt.Scan(&key)
  delete(bookmarks, key)
}
