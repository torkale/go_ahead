package main

import "fmt"

func (d *Dude) FullName() string {
  return fmt.Sprintf("%s %s", d.first, d.last)
}

func (d *Dude) SetFirst(name Name) {
  d.first = name
}

func (u *UberDude) ToString() string {
  return fmt.Sprintf("Name: %s Age: %d", u.FullName(), u.age)
}

func NewDude(f, l Name) *Dude {
  return &Dude{f, l}
}

func methods() {
  udi := new(UberDude)
  udi.first = "Sideshow"
  udi.last = "Bob"
  udi.age = 42
  fmt.Println(udi.ToString())
}
