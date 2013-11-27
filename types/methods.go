package main

import "fmt"

func (dude Person) FullName() string {
    return fmt.Sprintf("%s %s", dude.first, dude.last)
}

func (dude Person) SetFirst(name Name) {
    dude.first = name
}

func (h *Hero) ToString() string {
    return fmt.Sprintf("Name: %s Power: %s", h.FullName(), h.power)
}

func NewPerson(f, l Name) *Person {
    return &Person{f, l}
}

func methods() {
    superman := new(Hero)
    superman.first = "Klark"
    superman.last = "Kent"
    superman.power = "Flying"
    fmt.Println(superman.ToString())
}
