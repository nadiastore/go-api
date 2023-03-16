package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	// TODO: Buat reponya biar bisa diinstall
	"github.com/nadiastore/go-api/pkg/utils"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnection() (*sqlx.DB, error) {
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")

	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("mysql", mysqlConnURL)

	if err != nil {
		return nil, fmt.Errorf("failed connect to database, %w", err)
	}

	db.SetMaxOpenConns(maxConn)                           // default: 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // default: 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // default: 0 (re-use forever)

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("database unreachable, %w", err)
	}

	return db, nil
}
