package main

import "fmt"

var a string

func Init() {
  a = "finally started"
}

func doSomethingElse() {

}

func Simple() {
  go Init()
  doSomethingElse()
  fmt.Println(a)
}
