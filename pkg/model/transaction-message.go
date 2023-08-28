package model

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TransactionMessage struct {
	TransactionId string `json:"transactionId"`
}

func (tm *TransactionMessage) DeserializeMessage(d *amqp.Delivery) {
	err := json.Unmarshal([]byte(d.Body), &tm)
	if err != nil {
		fmt.Println("Erro ao deserializar JSON:", err)
		return
	}
}
