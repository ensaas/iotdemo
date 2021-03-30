/*

 Copyright 2021 Advantech
 Author: chienhsiang.chen@advantech.com.tw

*/
package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var (
	amqpsUri   string = ""
	routingKey string = "Hello"
	queue             = "helloQueue"
)

func main() {
	cfg := new(tls.Config)
	cfg.InsecureSkipVerify = true
	conn, err := amqp.DialTLS(amqpsUri, cfg)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = channel.QueueBind(
		queue.Name,  // queue name
		routingKey,  // routing key
		"amq.topic", // exchange
		false,       // no-wait
		nil,
	)
	if nil != err {
		failOnError(err, fmt.Sprintf("bind queue: %s routing Key: %s] to exchange Failed.", queue.Name, routingKey))
	}

	msgs, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
