package main

import (
  "fmt"

  "github.com/anaskhan96/soup"
)

func main() {
  url := "https://coinranking.com/"
  fmt.Println("Source :", url)
  resp, _ := soup.Get(url)

  doc := soup.HTMLParse(resp)
  title := doc.Find("a", "class", "coin-list__body__row").Find("span", "class", "coin-name").Text()

  fmt.Println("Crypto name :", title)
}