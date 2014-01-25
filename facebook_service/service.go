package main

import (
  "log"
  "net/http"
)

var conn Connection

func init() {
  conn = Connection{Uri: `amqp://guest:guest@localhost:5672`}
  err := conn.Setup()
  if err != nil {
    log.Fatalf("Failed to init amqp connection: %s", err)
  }
}

func main() {
  log.Println("Starting server on 8080")
  http.HandleFunc("/event", EventHandler)
  http.ListenAndServe(":8080", nil)
}
