package main

import (
	"os"

	"github.com/JoaoRafa19/codepix/application/grpc"
	"github.com/JoaoRafa19/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB


func main() {
	database = db.ConectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
	
}