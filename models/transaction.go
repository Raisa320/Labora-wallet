package models

import "time"

type Transaction struct {
	Id        int       `json:"id"`
	Amount    float64   `json:"amount" validate:"required"`
	Date      time.Time `json:"date"`
	DestinyId int       `json:"walletReceive" validate:"required"`
	SourceId  int       `json:"walletOrigin" validate:"required"`
	Type      int       `json:"type" validate:"required"`
}

type TransactionDTO struct {
	Id      int       `json:"id"`
	Amount  float64   `json:"amount"`
	Destiny string    `json:"destiny"`
	Source  string    `json:"source"`
	Type    string    `json:"type"`
	Date    time.Time `json:"date"`
}

func (transaction *TransactionDTO) GetType() {
	if transaction.Type == "1" {
		transaction.Type = "Deposit"
		return
	}
	transaction.Type = "Withdraw"
}
