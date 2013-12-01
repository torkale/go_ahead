package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Request from %s", r.URL)
}

func main() {
  fmt.Println("Starting server on 8080")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
