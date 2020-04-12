package db

import (
	"database/sql"
)

type DataSource interface {
	GetDatabase() *sql.DB
}
