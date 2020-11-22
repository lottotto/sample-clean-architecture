package infrastructure

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	connection *sql.DB
}

func NewDB() *DB {
	c := NewConfig()
	d := &DB{
		Host:     c.DB.Host,
		Username: c.DB.Username,
		Password: c.DB.Password,
		DBName:   c.DB.DBName,
	}
	return d
}

func (d *DB) Begin() (*sql.Tx, error) {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Username, d.Password, d.DBName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	d.connection = db
	defer d.connection.Close()
	return d.connection.Begin()
}

func (d *DB) Connect() *sql.DB {
	return d.connection
}
