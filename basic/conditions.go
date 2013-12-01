package main

import (
  "math/rand"
  "fmt"
)

func IsEven(x int) bool {
  var even bool
  if x%2 == 0 {
    even = true
  }

  if x%2 == 0 {
    even = true
  } else {
    even = false
  }

  if mod := x%2; mod == 0 {
    even = true
  }
  return even
}
func conditions() {
  r := rand.Intn(100)
  fmt.Printf("%d is even? %t\n", r, IsEven(r))
}
