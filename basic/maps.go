package main

import "fmt"

func maps() {
  m := make(map[string]int)
  m["Ten"] = 10
  fmt.Println(m)

  capitals := map[string]string{
    "Jerusalem": "Israel",
    "Paris": "France",
    "London": "UK",
  }
  fmt.Println(capitals)

  delete(capitals, "London")
  v, ok := capitals["London"]
  fmt.Println("The capital:", v, "Present?", ok)
}
