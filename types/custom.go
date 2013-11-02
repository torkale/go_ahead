package main

type Name string

type Dude struct {
  first Name
  last Name
}

type UberDude struct {
  Dude 
  age int
}

type Dudesville struct {
  dudes []Dude
}
