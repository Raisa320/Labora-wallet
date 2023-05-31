package postgres

import (
	"context"
	"database/sql"

	"github.com/raisa320/Labora-wallet/models"
)

type TransactionStorage struct {
}

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

func (repo *TransactionStorage) Create(ctx context.Context, transaction models.Transaction) (created *models.Transaction, err error) {
	return nil, nil
}
