package main

import "fmt"

type PersonAction func(some Person) Name

func (guy Person) DoublePersonAction(action PersonAction) string {
  return fmt.Sprintf("%s %s", action(guy), action(guy))
}

func DuplicateName(guy Person) string {
  return guy.DoublePersonAction(func(guy Person) Name {
    return guy.first
  })
}

func (crowd Crowd) ConcatPersonActions(action PersonAction) string {
  var result string
  for _, dude := range crowd.people {
    result = fmt.Sprintf("%s %s", result, action(dude))
  }
  return result
}

func AllLastNames(crowd Crowd) string {
  return crowd.ConcatPersonActions(func(dude Person) Name {
    return dude.last
  })
}

func functions() {
  crowd := Crowd{
    []Person{
      {"Bart", "Simpson"},
      {"Crusty", "Clown"},
      {"Barney", "Gumble"},
    }}
  fmt.Println(AllLastNames(crowd))
  fmt.Println(DuplicateName(crowd.people[0]))
}
