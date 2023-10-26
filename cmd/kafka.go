/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/JoaoRafa19/codepix/application/kafka"
	"github.com/JoaoRafa19/codepix/infrastructure/db"
	"github.com/spf13/cobra"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",

	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Produzindo mensagem")
		producer := kafka.NewKafkaProducer()
		database := db.ConectDB(os.Getenv("env"))
		producer.Publish("Ola consumer", "teste")

		go producer.DeliveryReport()

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

}
