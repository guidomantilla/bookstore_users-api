package impl

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDataSource struct {
	database *sql.DB
}

func NewMysqlDataSource(username string, password string, url string) *MysqlDataSource {

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)

	database, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	if err = database.Ping(); err != nil {
		panic(err)
	}

	return &MysqlDataSource{
		database: database,
	}
}

func (mysqlDataSource *MysqlDataSource) GetDatabase() *sql.DB {
	return mysqlDataSource.database
}
