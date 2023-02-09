package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func (k *KafkaConsumer) Consume() {
	configMap := ckafka.ConfigMap{
		"boostrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":         os.Getenv("KafkaConsumerGroupId"),
	}

	consumer, err := ckafka.NewConsumer(&configMap)
	if err != nil {
		log.Fatal("error consuming kafka message" + err.Error())
	}

	topics := os.Getenv("KafkaReadTopic")

	consumer.Subscribe(topics, nil)
	fmt.Println("Kafka consumer has been started")

	for {
		message, err := consumer.ReadMessage(-1)
		if err != nil {
			k.MsgChan <- message
		}

		fmt.Println(message.String())
	}

}
