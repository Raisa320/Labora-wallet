package repositories

import (
	"errors"

	"github.com/raisa320/Labora-wallet/models"
)

var ErrEntityNotExists error = errors.New("entity doesn't exists")
var ErrDuplicatedEntity error = errors.New("duplicated Entity")
var ErrInvalidEntityState error = errors.New("entity state is invalid")

type Wallet interface {
	GetAll() ([]models.Wallet, error)
	GetById(id int) (*models.Wallet, error)
	Create(wallet models.Wallet) (created *models.Wallet, err error)
	Update(wallet models.Wallet) (updated *models.Wallet, err error)
	Delete(id int) error
}

type Log interface {
	GetAll() ([]models.Log, error)
	GetById(id int) (*models.Log, error)
	Create(log models.Log) (created *models.Log, err error)
	Update(log models.Log) (updated *models.Log, err error)
	Delete(id int) error
}
