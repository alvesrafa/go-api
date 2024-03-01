package database

import (
	"database/sql"
	"fmt"

	"alvesrafa.dev/study/configs"
	_ "github.com/lib/pq" // postgres driver
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user =%s password=%s dbname =%s sslmode=disabled", conf.Host, conf.Port, conf.User, conf.Pass)

	conn, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err

}
