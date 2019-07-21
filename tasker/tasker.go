package main

import (
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	amqpHost := flag.String("amqp", "amqp://guest:guest@localhost:5672/", "enter amqp server")
	host := flag.String("host", "http://localhost:8080", "server to receive messages to queue")
	flag.Parse()
	conn, err := amqp.DialConfig(*amqpHost, amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, 2*time.Second)
		},
	})
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"basic", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
		fmt.Printf("Sent basic message on rabbitmq: %s\n", body)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
	})
	u, _ := url.Parse(*host)
	log.Fatal(http.ListenAndServe(u.Host, nil))

	forever := make(chan bool)

	<-forever

}
