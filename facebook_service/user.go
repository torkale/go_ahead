package main

import (
  "encoding/json"
  "io"
  "io/ioutil"
)

type Entry struct {
  ChangedFields []string `json:"changed_fields"`
}

type User struct {
  Entries *[]Entry `json:"entry"`
  Body    []byte
}

func NewUser(io io.Reader) (u *User, err error) {
  body, err := ioutil.ReadAll(io)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(body, &u)
  u.Body = body
  if err != nil {
    return nil, err
  }
  return u, nil
}

func (u *User) Publish(conn Connection) (err error) {
  for _, entry := range *u.Entries {
    for _, field := range entry.ChangedFields {
      err = conn.Publish(u.Body, field)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
