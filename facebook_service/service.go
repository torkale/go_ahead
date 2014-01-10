package main

import (
  "encoding/json"
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

func event_handler(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    log.Println(err, "read")
    w.WriteHeader(400)
    return
  }

  u, err := parseBody(body)
  if err != nil {
    log.Println(err, "parse")
    w.WriteHeader(400)
    return
  }

  err = publish_events(body, *u.Entries)
  if err != nil {
    log.Println(err, "publish")
    w.WriteHeader(500)
    return
  }

  w.WriteHeader(200)
}

func parseBody(body []byte) (u *User, err error) {
  err = json.Unmarshal(body, &u)
  if err != nil {
    return nil, err
  }
  return u, nil
}

func publish_events(body []byte, entries []Entry) (err error) {
  for _, entry := range entries {
    for _, field := range entry.Changed_fields {
      err = conn.Publish(body, field)
      if err != nil {
        return err
      }
    }
  }
  return nil
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
  http.HandleFunc("/event", event_handler)
  http.ListenAndServe(":8080", nil)
}
