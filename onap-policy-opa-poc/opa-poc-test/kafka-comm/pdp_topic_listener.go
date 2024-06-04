package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/open-policy-agent/opa/logging"
)

type KafkaConsumer struct {
	bootstrapServers string
	topic            string
	groupID          string
}

func NewKafkaConsumer(bootstrapServers, topic, groupID string) *KafkaConsumer {
	return &KafkaConsumer{
		bootstrapServers: bootstrapServers,
		topic:            topic,
		groupID:          groupID,
	}
}

func (kc *KafkaConsumer) HeartBeat() {
	for range time.Tick(time.Second * 10) {
		logging.New().Info("Listening for messages on topic policy-pdp-pap")
	}
}

func (kc *KafkaConsumer) Consume() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": kc.bootstrapServers,
		"group.id":          kc.groupID,
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Printf("Error creating consumer: %v\n", err)
		return
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{kc.topic}, nil)
	if err != nil {
		fmt.Printf("Error subscribing to topic: %v\n", err)
		return
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run == true {
		select {
		case sig := <-signals:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				logging.New().Info("Received message on topic %s: %s\n",
					*e.TopicPartition.Topic, string(e.Value))
			case kafka.Error:
				logging.New().Error("Error: %v\n", e)
				run = false
			}
		}
	}
}

func main() {
	bootstrapServers := "localhost:29092"
	topic := "policy-pdp-pap"
	groupID := "policy-pap-new"

	consumer := NewKafkaConsumer(bootstrapServers, topic, groupID)
	consumer.Consume()

	consumer.HeartBeat()
}
