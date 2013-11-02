package channel

import "fmt"

var (
  a string
  ready chan bool
)

func Init() {
  a = "finally started"
  ready <- true
}

func doSomethingElse() {

}

func Simple() {
  ready = make(chan bool)
  go Init()
  doSomethingElse()
  <-ready
  fmt.Println(a)
}
