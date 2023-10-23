package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionComfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []*Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull" `
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyToID        string   `gorm:"column:pix_key_to_id;type:uuid;not null" valid:"-"`
	Status            string   `json:"status" valid:"notnull" gorm:"type:varchar(20)"`
	Description       string   `json:"description" valid:"notnull" gorm:"type:varchar(255)" `
	CancelDescription string   `json:"cancel_description" valid:"notnull" gorm:"type:varchar(255)"`
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom:       accountFrom,
		Amount:            amount,
		PixKeyTo:          pixKeyTo,
		Status:            TransactionPending,
		Description:       description,
		CancelDescription: "",
	}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("the amount must be grater than 0")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		
		return fmt.Errorf("invalid status for the transaction: %v", t.Status)
	}

	if t.PixKeyTo.AccountId == t.AccountFrom.ID {
		return errors.New("invalid status for the transaction")

	}

	if err != nil {
		return err
	}
	return nil

}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel() error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Confirm(description string) error {
	t.Status = TransactionComfirmed
	t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()
	return err
}
