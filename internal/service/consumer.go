package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"golang.org/x/exp/slog"
)

type handler struct {
	*slog.Logger
}

func (handler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (handler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *handler) ConsumeClaim(_ sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var str strings.Builder
		str.WriteString(fmt.Sprint("Headers | ", msg.Headers, "\n"))
		str.WriteString(fmt.Sprint("Key | ", msg.Key, "\n"))
		str.WriteString(fmt.Sprint("Offset | ", msg.Offset, "\n"))
		str.WriteString(fmt.Sprint("Partition | ", msg.Partition, "\n"))
		str.WriteString(fmt.Sprint("Topic | ", msg.Topic, "\n"))
		str.WriteString(fmt.Sprint("Value | ", string(msg.Value), "\n"))
		str.WriteString(fmt.Sprint("Timestamp | ", msg.Timestamp.String(), "\n"))
		h.Info(str.String())
	}
	return nil
}

func RunConsumer(ctx context.Context) {
	group, err := sarama.NewConsumerGroup([]string{addr}, group, config)
	if err != nil {
		panic(err)
	}
	defer group.Close()

	go func() {
		for err := range group.Errors() {
			slog.ErrorCtx(ctx, "consumer errors", err)
		}
	}()

	topics := []string{topic}
	err = group.Consume(ctx, topics, &handler{
		slog.Default(),
	})
	if err != nil {
		panic(err)
	}
}
