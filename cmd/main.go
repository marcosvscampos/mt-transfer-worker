package main

import (
	"fmt"

	"github.com/marcosvscampos/mt-transfer-worker/pkg/consumer"
)

func main() {
	fmt.Println("Hello, mt-transfer-worker")

	consumer.StartConsumer()
	//client.CallAuthorizer()
}
