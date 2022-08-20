package service

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func RunProducer(msg string) {
	producer, err := sarama.NewAsyncProducer([]string{addr}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, producerErrors int
	for {
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(msg)}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("| Failed to produce message", err)
			producerErrors++
		case ok := <-producer.Successes():
			log.Println("| Success", ok)
			os.Exit(1)
		case <-signals:
			log.Printf("| Enqueued: %d; errors: %d\n", enqueued, producerErrors)
			return
		}
	}
}
