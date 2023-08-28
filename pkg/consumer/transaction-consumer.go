package consumer

import (
	"log"

	"github.com/marcosvscampos/mt-transfer-worker/pkg/model"
	"github.com/marcosvscampos/mt-transfer-worker/pkg/services"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func StartConsumer() {
	conn, err := amqp.Dial("amqps://qvliahds:oxVuUHwg_qAeFBI5CDPKHSEwbFjfAgG1@jackal.rmq.cloudamqp.com/qvliahds")

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		"queue.transactions", // queue
		"",                   // consumer
		false,                // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			tm := model.TransactionMessage{}
			tm.DeserializeMessage(&d)
			services.EffectTransfer(&tm)
			log.Println("Transaction ID:", tm.TransactionId)
			log.Printf("Done")

			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
