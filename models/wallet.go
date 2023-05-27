package models

import "time"

type Wallet struct {
	ID        string    `json:"id"`
	Person_id string    `json:"personId" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	Country   string    `json:"country"`
	Deleted   bool      `json:"deleted"`
}
