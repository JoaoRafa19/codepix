package usecase

import (
	"errors"

	"github.com/JoaoRafa19/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface

}

func (p* PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error){
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	pix,err := p.PixKeyRepository.RegisterKey(pixKey)

	if pix.ID == "" || err != nil {
		return nil, errors.New("unable to create a new key at the moment")
	}

	return pixKey, nil
}


func (p* PixUseCase) FindKey(key string, kind string) (*model.PixKey, error){
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

