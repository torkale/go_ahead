package main

import "fmt"

type Talker interface {
    Talk() string
}

func (dude Person) Talk() string {
    return fmt.Sprintf("My name is %s", dude.FullName())
}

func MakeSomeoneTalk(talker Talker) string {
    return talker.Talk()
}

func interfaces() {
    fmt.Println(MakeSomeoneTalk(NewPerson("Robert", "de Niro")))
}
