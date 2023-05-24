package models

import "time"

type Wallet struct {
	ID        string    `json:"id"`
	Person_id string    `json:"name" validate:"required"`
	Date      time.Time `json:"orderDate" validate:"required"`
	Country   string    `json:"country"`
}
