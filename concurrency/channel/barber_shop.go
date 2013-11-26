package channel

import (
    "fmt"
    "time"
)

const (
    SHAVE_TIME        = 700
    CLIENT_CYCLE_TIME = 400
    GREEN             = "\x1b[32;1m"
    YELLOW            = "\x1b[33;1m"
    RED               = "\x1b[31;1m"
)

type Customer struct {
    val int
}

var (
    seats = make(chan Customer, 2)
)

func barber() {
    for {
        c := <-seats
        time.Sleep(SHAVE_TIME * time.Millisecond)
        fmt.Println(GREEN, "   Barber is shaving customer", c.val)
    }
}

func (c Customer) enter() {
    select {
    case seats <- c:
    default:
        fmt.Println(RED, "<= No room for customer", c.val, "- Leaves shop")
    }
}

func customerProducer() {
    for i := 0; ; i++ {
        time.Sleep(CLIENT_CYCLE_TIME * time.Millisecond)
        c := Customer{i}
        fmt.Println(YELLOW, "=> Customer", c.val, "enters the shop")
        c.enter()
    }

}

func BarberShop() {
    go barber()
    go customerProducer()
    time.Sleep(8 * time.Second)
}
