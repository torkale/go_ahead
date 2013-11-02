package main
import "fmt"
func declarations() {
  var i int
  i = getInteger()
  j := getInteger()

  value, err := getValueOrError()

  value2, _ := getValueOrError()

  fmt.Println("Error:", err)

  fmt.Printf("Sum of declarations: %d\n", i + j + value + value2)
}

