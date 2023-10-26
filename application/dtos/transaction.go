package dtos

import (
	"encoding/json"
	"fmt"

	"github.com/asaskevich/govalidator"
)

type TransactionDTO struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error"`
}

func (t *TransactionDTO) isValid() error {

	v, err := govalidator.ValidateStruct(t)
	if err != nil && !v {
		fmt.Printf("error during transaction validation: %v\n", err.Error())
		return fmt.Errorf("error during transaction validation: %v", err.Error())
	}

	return nil

}

func (t *TransactionDTO) ParseJson(data []byte) error {

	if err := json.Unmarshal(data, t); err != nil {
		return err
	}

	if err := t.isValid(); err != nil {
		return err
	}

	return nil
}

func NewTransactionDTO() *TransactionDTO {
	return &TransactionDTO{}
}

func (t *TransactionDTO) ToJson() ([]byte, error ){

	if err:=t.isValid();err!=nil{
		return nil, err
	}

	result, err := json.Marshal(t);
	if err != nil {
		return nil, err
	}

	return result, nil

}