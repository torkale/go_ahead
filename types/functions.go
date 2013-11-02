package main

import "fmt"

type DudeAction func(dude Dude) Name

func (vil *Dudesville) ConcatDudeActions(action DudeAction) string {
  var result string
  for _, dude := range vil.dudes {
    result = fmt.Sprintf("%s %s", result, action(dude))
  }
  return result
}

func AllDudeLastNames(vil *Dudesville) string {
  return vil.ConcatDudeActions(func(d Dude) Name {
    return d.last
  })
}

func functions() {
  vil := Dudesville{
    []Dude{
      Dude{"Bart", "Simpson"},
      Dude{"Crusty", "Clown"},
    }}
  fmt.Println(AllDudeLastNames(&vil))
}
