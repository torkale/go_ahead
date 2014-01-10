package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

var conn Connection

type Entry struct {
  Changed_fields []string `json:"changed_fields"`
}

type User struct {
  Entries *[]Entry `json:"entry"`
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if retire("read", err, w) {
    return
  }

  var u *User
  err = json.Unmarshal(body, &u)
  if retire("parse", err, w) {
    return
  }

  for _, entry := range *u.Entries {
    for _, field := range entry.Changed_fields {
      err = conn.Publish(body, field)
      if retire("publish", err, w) {
        return
      }
    }
  }

  w.WriteHeader(200)
  fmt.Fprintf(w, `{"success":true}`)
}

func retire(phase string, err error, w http.ResponseWriter) bool {
  if err == nil {
    return false
  }
  log.Println(err, phase)
  w.WriteHeader(500)
  return true
}

func init() {
  conn = Connection{Uri: `amqp://guest:guest@localhost:5672`}
  err := conn.Setup()
  if err != nil {
    log.Fatalf("Failed to init amqp connection: %s", err)
  }
}

func main() {
  log.Println("Starting server on 8080")
  http.HandleFunc("/event", eventHandler)
  http.ListenAndServe(":8080", nil)
}
