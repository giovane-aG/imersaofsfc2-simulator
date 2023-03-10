package main

import (
	"log"

	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafka2 "github.com/giovane-aG/imersaofsfc2-simulator/application/kafka"
	"github.com/giovane-aG/imersaofsfc2-simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		go kafka2.Produce(msg)
		fmt.Println(string(msg.Value))
	}
}
