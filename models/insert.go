package models

import "github.com/kenfav/cong-reports-go/db"

func Insert(publisher Publisher) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO publishers (name, birthdate, baptismdate, othersheep) VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, publisher.Name, publisher.BirthDate, publisher.BaptismDate, publisher.OtherSheep).Scan(&id)
	return
}
