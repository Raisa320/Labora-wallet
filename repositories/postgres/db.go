package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/raisa320/Labora-wallet/repositories"
)

type DbConnection struct {
	*sql.DB //SeudoHerencia
}

var Db DbConnection

func InitDB() {
	err := connect_BD()
	if err != nil {
		log.Fatal(err)
	}
}

// PingOrDie envía un ping a la base de datos y si no se puede alcanzar, registra un error fatal.
func (db *DbConnection) PingOrDie() {
	if err := db.Ping(); err != nil {
		log.Fatalf("no se puede alcanzar la base de datos, error: %v", err)
	}
}

var dbConn *sql.DB

// Connect_BD conecta con la base de datos y devuelve un error si falla la conexión.
func connect_BD() error {

	var errDb error
	dbConfig, errDb := repositories.LoadEnvVariables()

	if errDb != nil {
		log.Fatalf("Error loading .env file: %s", errDb)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)
	var err error
	dbConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos:", dbConn)
	Db = DbConnection{dbConn}
	Db.PingOrDie()
	return err
}

// Interface for things that performs an scan over a given row.
// Actually it is a common interface for https://pkg.go.dev/database/sql#Rows.Scan and https://pkg.go.dev/database/sql#Row.Scan
type RowScanner interface {
	Scan(dest ...interface{}) error
}

// Scans a single row from a given query
type RowScanFunc func(rows RowScanner) (interface{}, error)

// Scans multiples rows using a scanner function in order to build a new "scanable" struct
func ScanMultiples(rows *sql.Rows, rowScanFunc RowScanFunc) ([]interface{}, error) {
	scaneables := []interface{}{}
	for rows.Next() {
		scanable, err := rowScanFunc(rows)
		if scanable == nil {
			return nil, err
		}
		scaneables = append(scaneables, scanable)
	}
	err := rows.Err()
	if err != nil {
		return nil, err
	}
	return scaneables, nil
}

func init() {
	InitDB()
	Db.PingOrDie()
}
