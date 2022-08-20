package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
)

type handler struct{}

func (handler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (handler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var str strings.Builder
		str.WriteString(fmt.Sprint("Headers | ", msg.Headers, "\n"))
		str.WriteString(fmt.Sprint("Key | ", msg.Key, "\n"))
		str.WriteString(fmt.Sprint("Offset | ", msg.Offset, "\n"))
		str.WriteString(fmt.Sprint("Partition | ", msg.Partition, "\n"))
		str.WriteString(fmt.Sprint("Topic | ", msg.Topic, "\n"))
		str.WriteString(fmt.Sprint("Value | ", string(msg.Value), "\n"))
		str.WriteString(fmt.Sprint("Timestamp | ", msg.Timestamp.String(), "\n"))
		fmt.Println(str.String())
	}
	return nil
}

func RunConsumer() {
	group, err := sarama.NewConsumerGroup([]string{addr}, group, config)
	if err != nil {
		panic(err)
	}
	defer group.Close()

	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	ctx := context.Background()
	for {
		topics := []string{topic}
		err := group.Consume(ctx, topics, new(handler))
		if err != nil {
			panic(err)
		}
	}
}
