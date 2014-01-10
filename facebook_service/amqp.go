package main

import (
  "github.com/streadway/amqp"
  "log"
  "time"
)

type Connection struct {
  conn *amqp.Connection
  pub  *amqp.Channel
  Uri  string
}

func (c *Connection) Setup() error {
  // Connects opens an AMQP connection from the credentials in the URL.
  connection, err := amqp.Dial(c.Uri)
  if err != nil {
    log.Println("Failed to open connection:", err)
    return err
  }

  channel, err := connection.Channel()
  if err != nil {
    log.Println("Failed to open channel:", err)
    return err
  }

  // We declare our topology on both the publisher and consumer to ensure they
  // are the same.  This is part of AMQP being a programmable messaging model.
  //
  err = channel.ExchangeDeclare("facebook_events", "topic", true, false, false, false, nil)
  if err != nil {
    log.Println("Failed to declare exchange:", err)
    return err
  }

  _, err = channel.QueueDeclare("location", true, false, false, false, nil)
  if err != nil {
    log.Println("Unable to declare queue", err)
    return err
  }

  err = channel.QueueBind("location", "location", "facebook_events", false, nil)
  if err != nil {
    log.Println("Unable to bind queue", err)
    return err
  }

  log.Println("Connection established")
  c.pub, c.conn = channel, connection
  return nil
}

func (c *Connection) Publish(body []byte, routingKey string) error {
  err := c.pub.Publish(
    "facebook_events", // publish to an exchange
    routingKey,        // routing to 0 or more queues
    false,             // mandatory
    false,             // immediate
    amqp.Publishing{
      ContentType:  "application/json",
      Body:         body,
      Timestamp:    time.Now(),
      DeliveryMode: amqp.Persistent, // 1=non-persistent, 2=persistent
      Priority:     0,               // 0-9
      // a bunch of application/implementation-specific fields
    })
  if err != nil {
    log.Println("publish", "Failed to publish: %s", err)
    return err
  }
  return nil
}
