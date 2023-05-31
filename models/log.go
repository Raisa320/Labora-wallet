package models

import "time"

type Log struct {
	ID        string    `json:"id"`
	Person_id string    `json:"national_id"`
	Date      time.Time `json:"date"`
	Status    bool      `json:"status"`
	Country   string    `json:"country"`
	Check_id  string    `json:"check_id"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
}

func (log *Log) SetStatus(score int) {
	if score != 1 {
		log.Status = false
	} else {
		log.Status = true
	}
}

func (log Log) GetStatus() string {
	if log.Status {
		return "Completado"
	}
	return "Rechazado"
}
