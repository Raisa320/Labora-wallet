package postgres

import (
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
		SELECT *
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

func (repo *WalletStorage) GetById(id int) (*models.Wallet, error) {
	row := Db.QueryRow(`
		SELECT *
		FROM wallet
		WHERE id = $1`, id)
	return scanWallet(row)
}

func (repo *WalletStorage) Create(animal models.Wallet) (*models.Wallet, error) {
	createQuery := `INSERT INTO animal (name, kind) VALUES ($1, $2) returning id`
	// err := Db.QueryRow(createQuery, animal.Name, animal.Kind).Scan(&animal.Id)
	err := Db.QueryRow(createQuery)
	if err != nil {
		return nil, err.Err()
	}
	return &animal, nil
}

func (repo *WalletStorage) Update(animal models.Wallet) (*models.Wallet, error) {
	// if animal.Id == nil {
	// 	return nil, repositories.ErrEntityNotExists
	// }
	updateQuery := `UPDATE animal SET name = $1, kind = $2 WHERE id = $3`
	// _, err := Db.Exec(updateQuery, animal.Name, animal.Kind, animal.Id)
	_, err := Db.Exec(updateQuery)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}

func (repo *WalletStorage) Delete(id int) (err error) {
	deleteQuery := `DELETE FROM animal WHERE id`
	_, err = Db.Exec(deleteQuery, id)
	return

}

// Scans a row interpreting it as 'models.Wallet' struct
func scanWallet(rows RowScanner) (*models.Wallet, error) {
	var wallet models.Wallet

	err := rows.Scan(&wallet.ID, &wallet.Person_id, &wallet.Date, &wallet.Country)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}
