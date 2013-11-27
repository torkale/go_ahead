package main

import (
    "fmt"
    "time"
)

func MeasureStart(label string) (string, time.Time) {
    return label, time.Now()
}

func Measure(label string, startTime time.Time) {
    duration := time.Now().Sub(startTime)
    fmt.Println(label, "Computation took", duration)
}

func benchmark() {
    defer Measure(MeasureStart("benchmark()"))
    time.Sleep(time.Second)
}
