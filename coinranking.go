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

  // Find the coins list, limited to top 20
  coins := doc.Find(".coin-list__body__row").Slice(0,20)
  var counter int = 0 

  coins.Each(func(i int, s *goquery.Selection) {
    // For each item found, get the name
    name := s.Find(".coin-name").Text()
    price := s.Find(".coin-list__body__row__price__value").Text()
    
    counter ++
    fmt.Printf("%v. %v $%v \n", counter, name, price)
  })
}

func main() {
  ScrapeForTop10Coins()
}