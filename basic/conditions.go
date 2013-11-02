package main

import "fmt"

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
  fmt.Printf("%d is even? %t\n", 5, IsEven(5))
}
