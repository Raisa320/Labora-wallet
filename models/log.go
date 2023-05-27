package models

import "time"

type Log struct {
	ID        string    `json:"id"`
	Person_id string    `json:"national_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
	Country   string    `json:"country"`
	Check_id  string    `json:"check_id"`
}
