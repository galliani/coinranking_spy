package main

import (
  "fmt"
  "log"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "github.com/fatih/color"
)

const baseURL string = "https://coinranking.com/"

func getColors(amount string, isNegative bool) string {
  if isNegative {

    red := color.New(color.FgRed).SprintFunc()
    return red(amount)

  } else {

    green := color.New(color.FgGreen).SprintFunc()
    return green(amount)
    
  }
}

func ScrapeForTop10Coins() {
  doc, err := goquery.NewDocument(baseURL)
  if err != nil {
    log.Fatal(err)
  }

  // Find the coins list, limited to top 20
  coins := doc.Find(".coin-list__body__row").Slice(0,20)
  var counter int = 0 

  coins.Each(func(i int, s *goquery.Selection) {
    var sign string
    // For each item found, get the name
    name := s.Find(".coin-name").Text()
    price := s.Find(".coin-list__body__row__price__value").Text()
    change := s.Find(".coin-list__body__row__change")
    isNegative := change.HasClass("coin-list__body__row__change--negative")
    
    // This is necessary because the original text parsed contains whitespaces
    amount := strings.TrimSpace(change.Text())    

    if isNegative { 
      sign = "-" 
    } else { 
      sign = "+" 
    }

    stringifiedAmount := sign + amount

    counter ++

    fmt.Printf("%v. %v $%v %v \n", counter, name, price, getColors(stringifiedAmount, isNegative))
  })
}

func main() {
  ScrapeForTop10Coins()
}