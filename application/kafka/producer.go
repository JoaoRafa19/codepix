package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	kafkaProducer *ckafka.Producer
	deliveryChan  chan ckafka.Event
}

func NewKafkaProducer() *Producer {

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}
	p, err := ckafka.NewProducer(&configMap)

	if err != nil {

		panic(err)
	}

	return &Producer{
		kafkaProducer: p,
		deliveryChan:  make(chan ckafka.Event),
	}
}

func (p *Producer) DeliveryReport() {
	for e := range p.deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery faild: ", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to: ", ev.TopicPartition)
			}
		}
	}
}

func (p *Producer) Publish(msg string, topic string) error {

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := p.kafkaProducer.Produce(message, p.deliveryChan)

	if err != nil {
		return err
	}
	return nil
}
