package main

import (
  "errors"
  "math/rand"
)

func main() {
  declarations()
  //conditions()
  //Factorial(2)
  //NextPowerOf2(1010)
  //slices()
  //maps()
}

func getInteger() int {
  return 10
}

func getValueOrError() (int, error) {
  if rand.Intn(10) < 5 {
    return 10, nil
  } else {
    return -1, errors.New("error")
  }
}
