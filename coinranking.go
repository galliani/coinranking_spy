package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "log"
)

const baseURL string = "https://coinranking.com/"

func ScrapeForTop10Coins() {
  doc, err := goquery.NewDocument(baseURL)
  if err != nil {
    log.Fatal(err)
  }

  // Find the coins list
  doc.Find(".coin-list__body__row").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the name
    name := s.Find(".coin-name").Text()
    fmt.Println("Crypto name :", name)
  })
}

func main() {
  ScrapeForTop10Coins()
}