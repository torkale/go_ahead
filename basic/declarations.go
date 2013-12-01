package main
import (
  "errors"
  "math/rand"
  "fmt"
)

func declarations() {
  var i int
  i = getInteger()
  j := getInteger()

  value, err := getValueOrError()

  value2, _ := getValueOrError()

  fmt.Println("Error:", err)

  fmt.Printf("Sum of declarations: %d\n", i + j + value + value2)
}

func getInteger() int {
  return 10
}

func getValueOrError() (int, error) {
  if rand.Intn(10) < 5 {
    return 1, nil
  } else {
    return -1, errors.New("random is too small")
  }
}
