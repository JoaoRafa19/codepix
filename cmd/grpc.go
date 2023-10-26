/*
Copyright © 2023  JoaoRafa19 / joaopedrorafael19@gmail.com
*/
package cmd

import (
	"os"

	"github.com/JoaoRafa19/codepix/application/grpc"
	"github.com/JoaoRafa19/codepix/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start a gRPC server",
	// Long: ``,
	Run: func(_ *cobra.Command, _ []string) {
		database := db.ConectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, grpcPortNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	grpcCmd.Flags().IntVarP(&grpcPortNumber, "port", "p", 50051, "gRPC server port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
