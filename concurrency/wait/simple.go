package wait

import (
  "sync"
  "fmt"
)

var (
  a string
  wg sync.WaitGroup
)

func Init() {
  defer wg.Done()
  a = "finally started"
}

func doSomethingElse() {

}

func Simple() {
  wg.Add(1)
  go Init()
  wg.Wait()
  doSomethingElse()
  fmt.Println(a)
}

