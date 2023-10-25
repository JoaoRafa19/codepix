package grpc

import (
	"context"

	"github.com/JoaoRafa19/codepix/application/grpc/pb"
	"github.com/JoaoRafa19/codepix/application/usecase"
)

type PixGrpcService struct {
	PixUsecase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUsecase.RegisterKey(in.Key, in.Kind, in.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     key.ID,
		Status: "created",
	}, nil

}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUsecase.FindKey(in.Key, in.Kind)
	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:        pixKey.ID,
		Kind:      pixKey.Kind,
		Key:       pixKey.Key,
		CraetedAt: pixKey.CreatedAt.String(),
		Account: &pb.Account{
			AccountId:     pixKey.AccountId,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
	}, nil
}


func NewgRPCService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUsecase: usecase,
	}
}