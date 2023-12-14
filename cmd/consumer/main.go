package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("consumer error: " + err.Error())
	}

	topics := []string{"teste"}
	consumer.SubscribeTopics(topics, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
