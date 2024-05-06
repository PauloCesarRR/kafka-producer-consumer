package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic:     "meu-topico",
		Key:       nil,
		Value:     sarama.StringEncoder("PauloCesarRR"),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Time{},
	}

	partition, offset, err := producer.SendMessage(message)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Mensagem enviada para o topico meu-topico, na partiçao %d, no offset %d \n", partition, offset)

}
