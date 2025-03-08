package core

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DBWrapperInterface permite abstraer la conexión a la BD.
type DBWrapperInterface interface {
	GetDB() *sql.DB
}

// DBWrapper es un envoltorio para la conexión a la base de datos.
type DBWrapper struct {
	DB *sql.DB
}

func (w *DBWrapper) GetDB() *sql.DB {
	return w.DB
}

func ConnectDB() (*DBWrapper, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DBWrapper{DB: db}, nil
}
