/*
Copyright © 2023  JoaoRafa19 / joaopedrorafael19@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/JoaoRafa19/codepix/application/grpc"
	"github.com/JoaoRafa19/codepix/application/kafka"
	"github.com/JoaoRafa19/codepix/infrastructure/db"
	"github.com/spf13/cobra"
)

var grpcPortNumber int

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Start gRPC and Kafka Consumer",

	Run: func(_ *cobra.Command, _ []string) {
		database := db.ConectDB(os.Getenv("env"))
		go grpc.StartGrpcServer(database, grpcPortNumber)

		fmt.Println("Produzindo mensagens...")
		producer := kafka.NewKafkaProducer()

		go producer.DeliveryReport()

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&grpcPortNumber, "grpc-port", "p", 50051, "gRPC server port")

}
