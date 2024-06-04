package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(bootstrapServers, topic string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{
		producer: p,
		topic:    topic,
	}, nil
}

func (kp *KafkaProducer) Produce(message string) error {
	deliveryChan := make(chan kafka.Event)

	err := kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}

func (kp *KafkaProducer) Close() {
	kp.producer.Close()
}

func main() {
	bootstrapServers := "localhost:29092"
	topic := "policy-pdp-pap"

	producer, err := NewKafkaProducer(bootstrapServers, topic)
	if err != nil {
		fmt.Printf("Error creating Kafka producer: %v\n", err)
		return
	}
	defer producer.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			message := []byte(fmt.Sprint(`{"pdpType":"opa","state":"ACTIVE","healthy":"HEALTHY","description":"Pdp Heartbeat","messageName":"PDP_STATUS","requestId":"523b3bc5-2e56-5054-b5eg-36f5ba130513""pdpGroup":"defaultGroup","pdpSubgroup":"opa"}`))
			err := producer.Produce(string(message))
			if err != nil {
				fmt.Printf("Error producing message: %v\n", err)
			} else {
				fmt.Println("Message sent successfully")
			}
		case <-quit:
			fmt.Println("Received termination signal. Exiting...")
			return
		}
	}
}
