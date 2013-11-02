package channel

import (
  "time"
)

func ProducerConsumerTimeout() {
  c := make(chan string)
  go producer(c)

  abort := time.After(2*time.Second)
  consumer_with_abort(c, abort)
}

func consumer_with_abort(c chan string, abort <-chan time.Time) {
  for {
    select {
    case msg := <-c:
      process(msg)
    case <- abort:
      return
    }
  }
}

