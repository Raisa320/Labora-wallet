package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/go-playground/validator"
	"github.com/raisa320/Labora-wallet/models"
)

type TransactionStorage struct {
}

var mutex sync.Mutex

func NewTransactionStorage() *TransactionStorage {
	return &TransactionStorage{}
}

func (repo *TransactionStorage) GetById(ctx context.Context, id int) (*models.TransactionDTO, error) {
	var transaction models.TransactionDTO
	getQuery := `
	SELECT t.id,t.amount,t.type, source.person_name as "source", destiny.person_name as "destiny"
	FROM transaction as t 
	INNER JOIN wallet as  source ON source.id = t.source_id
	INNER JOIN wallet as destiny ON destiny.id=t.destiny_id  
	WHERE t.id=$1;`
	err := Db.QueryRowContext(ctx, getQuery, id).Scan(&transaction.Id, &transaction.Amount, &transaction.Type, &transaction.Source, &transaction.Destiny)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	transaction.GetType()
	return &transaction, nil
}

func (repo *TransactionStorage) GetByWallet(walletId int) ([]models.TransactionDTO, error) {
	getQuery := `
	SELECT t.id,t.amount,t.type, t.date,source.person_name as "source", destiny.person_name as "destiny"
	FROM transaction as t 
	INNER JOIN wallet as  source ON source.id = t.source_id
	INNER JOIN wallet as destiny ON destiny.id=t.destiny_id  
	WHERE t.source_id=$1;`

	rows, err := Db.Query(getQuery, walletId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	transactions := []models.TransactionDTO{}
	for rows.Next() {
		var transaction models.TransactionDTO
		err := rows.Scan(&transaction.Id, &transaction.Amount, &transaction.Type, &transaction.Date, &transaction.Source, &transaction.Destiny)
		if err != nil {
			return nil, err
		}
		transaction.GetType()
		transactions = append(transactions, transaction)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (repo *TransactionStorage) Create(ctx context.Context, transaction models.Transaction) error {
	mutex.Lock()
	validate := validator.New()

	err := validate.Struct(transaction)
	if err != nil {
		return err
	}
	if transaction.Type <= 0 || transaction.Type > 2 {
		return fmt.Errorf("not valid type of transaction")
	}

	tx, err := Db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		mutex.Unlock()
	}()

	walletSource, err := NewWalletStorage().GetById(transaction.SourceId)
	if err != nil {
		return err
	}

	walletDestiny, err := NewWalletStorage().GetById(transaction.DestinyId)
	if err != nil {
		return err
	}

	if walletSource.Amount < transaction.Amount {
		return fmt.Errorf("insufficient funds to process the withdrawal: %v current balance", walletSource.Amount)
	}

	if transaction.Type == 2 && walletSource.ID != walletDestiny.ID {
		return fmt.Errorf("it is not a valid operation")
	}

	amoutSource := walletSource.Amount - transaction.Amount
	amoutDestiny := walletDestiny.Amount + transaction.Amount
	err = NewWalletStorage().UpdateAmount(amoutSource, *walletSource, tx)
	if err != nil {
		return err
	}
	err = NewWalletStorage().UpdateAmount(amoutDestiny, *walletDestiny, tx)
	if err != nil {
		return err
	}

	createQuery := `INSERT INTO transaction(
	amount, destiny_id, source_id, type)
	VALUES ($1, $2, $3, $4) returning id`
	err = tx.QueryRowContext(ctx, createQuery, transaction.Amount, transaction.DestinyId, transaction.SourceId, transaction.Type).Scan(&transaction.Id)
	if err != nil {
		return err
	}
	return nil
}
