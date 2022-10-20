package models

import "github.com/kenfav/cong-reports-go/db"

func Update(id int64, publisher Publisher) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE publishers SET name=$1, birthdate=$2, baptismdate=$3, othersheep=$4 WHERE id=$5`,
		publisher.Name, publisher.BirthDate, publisher.BaptismDate, publisher.OtherSheep, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
