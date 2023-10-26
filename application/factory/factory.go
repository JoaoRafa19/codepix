package factory

import (
	"github.com/JoaoRafa19/codepix/application/usecase"
	"github.com/JoaoRafa19/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUsecaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository:= repository.PixKeyRepositoryDB{Db: database}
	transactionRepository:= repository.TransactionRepositoryDB{Db: database}
	
	transactionUsecase := usecase.TransactionUseCase{
		TransactionRepository: transactionRepository,
		PixRepository: pixRepository,
	}

	return transactionUsecase
}