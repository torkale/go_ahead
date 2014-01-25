package main

import (
  "fmt"
  "log"
  "net/http"
)

func EventHandler(w http.ResponseWriter, r *http.Request) {
  user, err := NewUser(r.Body)
  if err != nil {
    errorResponse("NewUser", err, w)
    return
  }

  err = user.Publish(conn)
  if err != nil {
    errorResponse("User.Publish", err, w)
    return
  }

  w.WriteHeader(200)
  fmt.Fprintf(w, `{"success":true}`)
}

func errorResponse(phase string, err error, w http.ResponseWriter) {
  log.Println(phase, err)
  w.WriteHeader(500)
  fmt.Fprintf(w, `{"success":false}`)
}
