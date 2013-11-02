package channel

import(
  "fmt"
  "sync"
)

type Customer string

var (
  seats = make(chan Customer, 2)
  wg = new(sync.WaitGroup)
  customers = []Customer{ "A", "B", "C", "D" }
)

func barber() {
  for {
    c := <-seats
    fmt.Println("Barber shaving", c)
  }
}

func (c Customer) enter() {
  defer wg.Done()
  select {
  case seats <- c:
  default:
    fmt.Println("Customer", c, "Leaves")
  }
}

func BarberShop() {
  wg.Add(4)
  go barber()
  for _, c := range customers { go c.enter() }
  wg.Wait()
}
