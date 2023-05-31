package repositories

import (
	"context"
	"errors"

	"github.com/raisa320/Labora-wallet/models"
)

var ErrEntityNotExists error = errors.New("entity doesn't exists")
var ErrDuplicatedEntity error = errors.New("duplicated Entity")
var ErrInvalidEntityState error = errors.New("entity state is invalid")

type Wallet interface {
	GetAll() ([]models.Wallet, error)
	GetById(id int) (*models.WalletDTO, error)
	Create(ctx context.Context, wallet models.Wallet) (created *models.Wallet, err error)
	Update(ctx context.Context, id int, wallet models.Wallet) (updated *models.Wallet, err error)
	Delete(id int) (success bool, err error)
}

type Log interface {
	GetAll() ([]models.Log, error)
	GetById(id int) (*models.Log, error)
	GetByPersonId(personId string) (*models.Log, error)
	Create(ctx context.Context, log models.Log) (created *models.Log, err error)
}

type Transaction interface {
	GetById(ctx context.Context, id int) (*models.TransactionDTO, error)
	GetByWallet(walletId int) ([]models.TransactionDTO, error)
	Create(ctx context.Context, transaction models.Transaction) (created *models.Transaction, err error)
}
