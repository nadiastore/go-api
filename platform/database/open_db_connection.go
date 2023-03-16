package database

import (
	"os"

	// TODO: Buat reponya biar bisa diinstall
	"github.com/nadiastore/go-api/app/queries"

	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.UserQueries // load query dari User model
	*queries.BookQueries // load query dari Book model
}

func OpenDBConnection() (*Queries, error) {
	var (
		db  *sqlx.DB
		err error
	)

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "pgx":
		db, err = PostgreSQLConnection()

	case "mysql":
		db, err = MysqlConnection()
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set query dari model:
		UserQueries: &queries.UserQueries{DB: db}, // dari User model
		BookQueries: &queries.BookQueries{DB: db}, // dari Book model
	}, nil
}
