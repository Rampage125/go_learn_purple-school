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
  fmt.Println(`
— 1. Посмотреть закладки
— 2. Добавить закладку
— 3. Удалить закладку
— 4. Выход`)
  fmt.Scan(&x)
  return x
}

func enterInfo(prompt string) (info string) {
  fmt.Println(prompt)
  fmt.Scan(&info)
  return
}

func addBookmark(bookmarks bookMap) {
  key := enterInfo("Введите ключ: ")
  value := enterInfo("Введите значение: ")
  bookmarks[key] = value
}

func deleteBookmark(bookmarks bookMap) {
  key := enterInfo("Введите ключ: ")
  delete(bookmarks, key)
}
