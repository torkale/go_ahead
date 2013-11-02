package main

import "fmt"

type Talker interface {
  Talk() string
}

func (d *Dude) Talk() string {
  return fmt.Sprintf("My name is %s", d.FullName())
}

func MakeSomeoneTalk(talker Talker) string {
  return talker.Talk()
}

func interfaces() {
  fmt.Println(MakeSomeoneTalk(&Dude{"Robert", "de Niro"}))
}
