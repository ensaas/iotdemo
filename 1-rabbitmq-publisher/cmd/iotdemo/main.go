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
	message  string = "Hello World!!!"
	amqpsUri string = ""
	topic    string = "Hello"
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

	err = channel.ExchangeDeclare(
		"amq.topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		failOnError(err, "Failed to declare the Exchange")
	}

	err = channel.Publish(
		"amq.topic",
		topic,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         []byte(message),
		})
	if err != nil {
		failOnError(err, "Failed to publish a message")
	} else {
		fmt.Println("Publish success.")
	}

}
