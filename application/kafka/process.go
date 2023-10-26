package kafka

import (
	"fmt"
	"os"

	"github.com/JoaoRafa19/codepix/application/dtos"
	"github.com/JoaoRafa19/codepix/application/factory"
	"github.com/JoaoRafa19/codepix/application/usecase"
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
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{os.Getenv("kafkaTransactionTopic"), os.Getenv("kafkaTransactionConfirmationTopic")}
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

func (k *KafkaProcessor) processMessage(msg *ckafka.Message) {
	transactionsTopic := "transactions"
	transacionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionsTopic:
		k.processTransaction(msg)
	case transacionConfirmationTopic:
		k.processTransactionConfirmation(msg)
	default:
		fmt.Println("not a valid topic", string(msg.Value))
		fmt.Println(topic)
	}
}

func (k *KafkaProcessor) processTransaction(msg *ckafka.Message) error {
	transactionDTO := dtos.NewTransactionDTO()

	err := transactionDTO.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	transactionUsecase := factory.TransactionUsecaseFactory(k.Database)

	createdTransaction, err := transactionUsecase.Register(
		transactionDTO.AccountID,
		transactionDTO.Amount,
		transactionDTO.PixKeyTo,
		transactionDTO.PixKeyKindTo,
		transactionDTO.Description,
	)

	if err != nil {
		fmt.Println("error registring transaction:", err)
		return err
	}

	topic := "bank" + createdTransaction.PixKeyTo.Account.Bank.Code
	transactionDTO.ID = createdTransaction.ID
	transactionDTO.Status = model.TransactionPending
	transactionJson, err := transactionDTO.ToJson()

	if err != nil {
		return err
	}

	err = k.Producer.Publish(string(transactionJson), topic)

	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaProcessor) processTransactionConfirmation(msg *ckafka.Message) error {

	transactionDTO := dtos.NewTransactionDTO()
	err := transactionDTO.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	transactionUsecase := factory.TransactionUsecaseFactory(k.Database)

	if transactionDTO.Status == model.TransactionComfirmed {
		err = k.confirmTransaction(transactionDTO, transactionUsecase)
		if err != nil {
			return err
		}
	} else if transactionDTO.Status == model.TransactionCompleted {
		_, err := transactionUsecase.Complete(transactionDTO.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *KafkaProcessor) confirmTransaction(transaction *dtos.TransactionDTO, usecase usecase.TransactionUseCase) error {
	confirmedTransaction, err := usecase.Confirm(transaction.ID)

	if err != nil {
		return err
	}

	topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
	transactionJson, err := transaction.ToJson()

	if err != nil {
		return err
	}

	err = k.Producer.Publish(string(transactionJson), topic)
	if err != nil {
		return err
	}

	return nil
}
