package channel

import (
  "fmt"
  "time"
)

func producer(c chan string){
  defer close(c)
  for {
    work := getWork()
    c <- work
  }
}

func consumer(c chan string) {
  for msg := range c {
      process(msg)
  }
}

func getWork() string{
  time.Sleep(100 * time.Millisecond)
  return fmt.Sprintf("Current time: %s", time.Now())
}

func process(message string) {
  fmt.Println(message)
}

func ProducerConsumer() {
  c := make(chan string)
  go producer(c)
  consumer(c)
}
