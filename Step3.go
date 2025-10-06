package main

import "fmt"


func main() {
  transactions := []float64{}
  for {
    trans := addTransaction()
    if trans == 0 {
      break
    } else {
      transactions = append(transactions, trans)
    }
  }
  res := 0.0
  for _,v := range transactions{
    res += v
  }
  fmt.Println("Баланс: \n",res)
}


func addTransaction() float64 {
  fmt.Println("Введите транзакцию: ")
  var trans float64
  fmt.Scan(&trans)
  return trans
}