package service

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

const (
	hostVar  = "KAFKA_HOST"
	portVar  = "KAFKA_PORT"
	topicVar = "KAFKA_TOPIC"
	groupVar = "KAFKA_CONSUMER_GROUP"
)

var (
	config *sarama.Config
	addr   string
	topic  string
	group  string
)

func RunConfig() {
	config = sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	config.Consumer.Return.Errors = true
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true

	addr = fmt.Sprintf("%s:%s", viper.GetString(hostVar), viper.GetString(portVar))
	topic = viper.GetString(topicVar)
	group = viper.GetString(groupVar)
}
