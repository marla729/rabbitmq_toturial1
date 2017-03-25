package main

import (
  "fmt"
  "log"

  "github.com/streadway/amqp"
)
//helper function to check the return value for each amqp call
func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
} 

func main(){
//connect to RabbitMQ server
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

//create a channel
ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()

//declare a queue
q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")

body := "hello world this is the first message" //guess: this is the message to be sent to, i.e. the message receiver will receive.
err = ch.Publish(
  "",     // exchange
  q.Name, // routing key
  false,  // mandatory
  false,  // immediate
  amqp.Publishing {
    ContentType: "text/plain",
    Body:        []byte(body),
  })
failOnError(err, "Failed to publish a message")
}
