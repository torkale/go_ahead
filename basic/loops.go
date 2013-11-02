package main

import "fmt"

func Factorial(num int) {
  factorial := 1
  for i := 2; i <= num; i++ {
    factorial *= i
  }
  fmt.Printf("Factorial(%d)==%d\n", num, factorial)
}

func NextPowerOf2(num int) {
  next := 1
  for next < num {
    next *=2
  }
  fmt.Printf("NextPowerOf2(%d)==%d\n", num, next)
}
