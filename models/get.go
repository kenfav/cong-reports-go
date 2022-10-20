package models

import "github.com/kenfav/cong-reports-go/db"

func Get(id int64) (publisher Publisher, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM publishers WHERE id=$1`, id)
	err = row.Scan(&publisher.ID, &publisher.Name, &publisher.BirthDate, &publisher.BaptismDate, &publisher.OtherSheep)
	return
}
