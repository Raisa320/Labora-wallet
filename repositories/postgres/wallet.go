package postgres

import (
	"context"
	"database/sql"

	"github.com/raisa320/Labora-wallet/models"
)

type WalletStorage struct {
}

func NewWalletStorage() *WalletStorage {
	return &WalletStorage{}
}

// Gets all legal entities for a list of organizations
func (repo *WalletStorage) GetAll() ([]models.Wallet, error) {
	rows, err := Db.Query(`
		SELECT id, person_id, date, country, amount
		FROM wallet`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	wallets := []models.Wallet{}
	for rows.Next() {
		wallet, err := scanWallet(rows)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, *wallet)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return wallets, nil
}

func (repo *WalletStorage) GetById(id int) (*models.WalletDTO, error) {
	row := Db.QueryRow(`
		SELECT id, person_name, amount
		FROM wallet
		WHERE id = $1`, id)
	var wallet models.WalletDTO
	err := row.Scan(&wallet.ID, &wallet.PersonName, &wallet.Amount)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	transactions, err := NewTransactionStorage().GetByWallet(id)
	if err != nil {
		return nil, err
	}
	wallet.Transaction = transactions
	return &wallet, err
}

func (repo *WalletStorage) Create(ctx context.Context, wallet models.Wallet) (*models.Wallet, error) {
	createQuery := `INSERT INTO wallet(
		person_id, date, country, person_name)
		VALUES ($1, $2, $3, $4) returning id`
	err := Db.QueryRowContext(ctx, createQuery, wallet.Person_id, wallet.Date, wallet.Country, wallet.PersonName).Scan(&wallet.ID)

	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (repo *WalletStorage) Update(ctx context.Context, id int, wallet models.Wallet) (*models.Wallet, error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, err
	}

	updateQuery := `UPDATE wallet SET have_card=$1  WHERE id = $2`
	_, err = tx.ExecContext(ctx, updateQuery, wallet.HavePhyscalCard, id)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró ningún objeto
			tx.Commit()
			return nil, nil
		}
		tx.Rollback()
		return nil, err
	}
	selectQuery := `SELECT id, person_id, date, country, amount, have_card FROM wallet	WHERE id = $1`
	err = tx.QueryRowContext(ctx, selectQuery, id).Scan(&wallet.ID, &wallet.Person_id, &wallet.Date, &wallet.Country, &wallet.Amount, &wallet.HavePhyscalCard)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &wallet, nil
}

func (repo *WalletStorage) Delete(id int) (success bool, err error) {

	deleteQuery := `DELETE FROM wallet WHERE id = $1`
	result, err := Db.Exec(deleteQuery, id)
	if err != nil {
		return false, err
	}
	numberRowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if numberRowsAffected == 0 {
		return false, nil
	}
	return true, nil

}

// Scans a row interpreting it as 'models.Wallet' struct
func scanWallet(rows RowScanner) (*models.Wallet, error) {
	var wallet models.Wallet

	err := rows.Scan(&wallet.ID, &wallet.Person_id, &wallet.Date, &wallet.Country, &wallet.Amount)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
