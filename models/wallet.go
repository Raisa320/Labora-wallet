package models

import "time"

type Wallet struct {
	ID              string    `json:"id"`
	Person_id       string    `json:"personId" validate:"required"`
	PersonName      string    `json:"personName"`
	Date            time.Time `json:"date" validate:"required"`
	Country         string    `json:"country"`
	Amount          float64   `json:"amount"`
	HavePhyscalCard bool      `json:"haveCard"`
}

type WalletDTO struct {
	ID          string           `json:"id"`
	PersonName  string           `json:"personName"`
	Amount      float64          `json:"amount"`
	Transaction []TransactionDTO `json:"movements"`
}
