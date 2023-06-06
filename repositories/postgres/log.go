package postgres

import (
	"context"
	"database/sql"

	"github.com/raisa320/Labora-wallet/models"
)

type LogStorage struct {
}

func NewLogStorage() *LogStorage {
	return &LogStorage{}
}

func (repo *LogStorage) GetAll() ([]models.Log, error) {
	rows, err := Db.Query(`
		SELECT *
		FROM log`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	logs := []models.Log{}
	for rows.Next() {
		log, err := scanLog(rows)
		if err != nil {
			return nil, err
		}
		logs = append(logs, *log)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (repo *LogStorage) GetById(id int) (*models.Log, error) {
	row := Db.QueryRow(`
		SELECT *
		FROM log
		WHERE id = $1`, id)
	return scanLog(row)
}

func (repo *LogStorage) GetByPersonId(personId string) (*models.Log, error) {
	row := Db.QueryRow(`
		SELECT *
		FROM log
		WHERE person_id = $1`, personId)
	log, err := scanLog(row)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return log, err
}

func (repo *LogStorage) Create(ctx context.Context, log models.Log) (*models.Log, error) {
	createQuery := `INSERT INTO log(
		person_id, date, status, country, check_id, type, message)
		VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`
	row := Db.QueryRowContext(
		ctx, createQuery, log.Person_id, log.Date, log.Status, log.Country, log.Check_id, log.Type, log.Message)
	err := row.Scan(&log.ID)

	if err != nil {
		return nil, err
	}
	return &log, nil
}

func scanLog(rows RowScanner) (*models.Log, error) {
	var log models.Log

	err := rows.Scan(&log.ID, &log.Person_id, &log.Date, &log.Status, &log.Country, &log.Check_id)
	if err != nil {
		return nil, err
	}
	return &log, nil
}
