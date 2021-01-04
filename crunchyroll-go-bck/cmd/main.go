package main

import (
  "fmt"
  "github.com/gokel166/anime-service/testutls"
)

func main() {
  b := "Get this string"
  c := printers.ReadTestString(b)

  secString := printers.TestPrinter()
  fmt.Println(c)
  fmt.Println(secString)
  fmt.Println("Here")
}

