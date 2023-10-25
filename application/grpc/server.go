package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/JoaoRafa19/codepix/application/grpc/pb"
	"github.com/JoaoRafa19/codepix/application/usecase"
	"github.com/JoaoRafa19/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pixRepository := repository.PixKeyRepositoryDB{Db: database}
	pixUsecase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewgRPCService(pixUsecase)
	
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)


	addres := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", addres)
	if err != nil {
		log.Fatalf("Not start gRPC server:%v", err)
	}

	log.Printf("grpc server start on port:%d", port)
	errors := grpcServer.Serve(listener)

	if errors != nil {
		log.Fatalf("Not start gRPC server:%v", err)
	}

}
