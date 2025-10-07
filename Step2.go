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

func enterKey() string {
  var key string
  fmt.Println("Введите ключ: ")
  fmt.Scan(&key)
  return key
}

func enterValue() string {
  var value string
  fmt.Println("Введите значение: ")
  fmt.Scan(&value)
  return value
}

func addBookmark(bookmarks bookMap) {
  key := enterKey()
  value := enterValue()
  bookmarks[key] = value
}

func deleteBookmark(bookmarks bookMap) {key := enterKey()
  delete(bookmarks, key)
}
