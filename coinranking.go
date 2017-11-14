package main

import (
  "fmt"
  "github.com/anaskhan96/soup"
  "os"
)

func main() {
  url := "https://coinranking.com/"
  fmt.Println("Source :", url)
  response, error := soup.Get(url)

  // Add error handling, will terminate everything if error is not nil
  if error != nil {
    os.Exit(1)
  }

  doc := soup.HTMLParse(response)
  title := doc.Find("a", "class", "coin-list__body__row").Find("span", "class", "coin-name").Text()

  fmt.Println("Crypto name :", title)
}