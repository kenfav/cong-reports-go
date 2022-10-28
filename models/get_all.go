package models

import "github.com/kenfav/cong-reports-go/db"

func GetAll() (publishers []Publisher, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM publishers`)
	if err != nil {
		return
	}
	for rows.Next() {
		var publisher Publisher

		err = rows.Scan(&publisher.ID, &publisher.Surname, &publisher.Name, &publisher.BirthDate, &publisher.BaptismDate, &publisher.OtherSheep, &publisher.Elder, &publisher.MinisterialServant, &publisher.RegularPionner)
		if err != nil {
			continue
		}
		publishers = append(publishers, publisher)
	}
	return
}
