package kafka

import (
	"fmt"

	"github.com/JoaoRafa19/codepix/domain/model"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

type KafkaProcessor struct {
	Database *gorm.DB
	Producer *Producer
}

func NewKafkaProcessor(database *gorm.DB, producer *Producer) *KafkaProcessor {
	return &KafkaProcessor{
		Database: database,
		Producer: producer,
	}
}

func (k *KafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "consumergroup",
		"auto.offset.reset": "earliest",
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)
	fmt.Println("kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.processMessage(msg)
			fmt.Println(string(msg.Value))
		}
	}

}


func (k *KafkaProcessor) processMessage (msg *ckafka.Message){
	transactionsTopic := "transactions"
	transacionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionsTopic:
	case transacionConfirmationTopic:
	default:
		fmt.Println("not a valid topic", string(msg.Value))
		fmt.Println(topic)
	}
}

func (k* KafkaProcessor) processTransaction(msg *ckafka.Message) error {
	transaction := model.NewTransaction()
}