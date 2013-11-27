package main

type Name string

type Person struct {
    first Name
    last  Name
}

type Hero struct {
    Person
    power string
}

type Crowd struct {
    people []Person
}
