package service

import (
	"github.com/Shopify/sarama"
	"golang.org/x/exp/slog"
)

func RunProducer(msg string) {
	producer, err := sarama.NewAsyncProducer([]string{addr}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	select {
	case producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(msg)}:
		slog.Info("producer", "message", msg)
	case err := <-producer.Errors():
		slog.Error("producer", "message failed", err)
	case ok := <-producer.Successes():
		slog.Info("producer", "success", ok)
	}
}
