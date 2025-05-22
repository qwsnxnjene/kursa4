package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	_ "modernc.org/sqlite"
)

var dbPath string
var Database *sql.DB

func getDbDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

// CreateDatabase получает путь к файлу с базой данных и, если необходимо, создаёт таблицу
func CreateDatabase() error {
	dbDir := getDbDir()
	dbPath = filepath.Join(dbDir, "unis.db")

	_, err := os.Stat(dbPath)
	var install bool

	if err != nil {
		install = true
	}

	if install {
		err := CreateTable(dbPath)
		if err != nil {
			return fmt.Errorf("[internal.app.db.CreateDatabase]: %w", err)
		}
	}

	return nil
}

// CreateTable создаёт таблицу в файле, указанном в path
// возвращает nil при успешном создании таблицы, иначе ошибку
func CreateTable(pathDb string) error {
	db, err := sql.Open("sqlite", pathDb)
	if err != nil {
		return fmt.Errorf("[internal.app.db.CreateTable]: %w", err)
	}

	dbDir := getDbDir()
	sqlPath := filepath.Join(dbDir, "unis.sql")
	createQuery, err := os.ReadFile(sqlPath)
	if err != nil {
		return fmt.Errorf("[internal.app.db.CreateTable]: %w", err)
	}

	query := string(createQuery)

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("[internal.app.db.CreateTable]: %w", err)
	}

	log.Printf("[internal.app.db.CreateTable] Created database table: %s", pathDb)
	return nil
}

// OpenSql возвращает таблицу для работы с ней
func OpenSql() (*sql.DB, error) {
	err := CreateDatabase()
	if err != nil {
		return nil, fmt.Errorf("[internal.app.db.OpenSql]: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("[internal.app.db.OpenSql]: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("[internal.app.db.OpenSql]: failed to ping database: %w", err)
	}

	Database = db
	return db, nil
}
