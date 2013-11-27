package main

import "fmt"

type PersonAction func(some Person) Name

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
        }}
    fmt.Println(AllLastNames(crowd))
}
