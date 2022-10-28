package models

import "github.com/kenfav/cong-reports-go/db"

func Insert(publisher Publisher) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO publishers (surname, name, birthdate, baptismdate, othersheep, elder, ministerialservant, regularpionner) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err = conn.QueryRow(sql, publisher.Surname, publisher.Name, publisher.BirthDate, publisher.BaptismDate, publisher.OtherSheep, publisher.Elder, publisher.MinisterialServant, publisher.RegularPionner).Scan(&id)
	return
}
