package db

import (
	"database/sql"
	"fmt"

	"github.com/kenfav/cong-reports-go/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s pass=%s dbname=%s sslmode=disabled",
		conf.Host, conf.Port, conf.User, conf.Database, conf.Pass, conf.Database)
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err

}
