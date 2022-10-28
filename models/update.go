package models

import "github.com/kenfav/cong-reports-go/db"

func Update(id int64, publisher Publisher) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE publishers SET surname=$2, name=$3, birthdate=$4, baptismdate=$5, othersheep=$6, elder=$7, ministerialservant=$8, regularpionner=$9, WHERE id=$1`,
		id, &publisher.Surname, &publisher.Name, &publisher.BirthDate,
    &publisher.BaptismDate, &publisher.OtherSheep, &publisher.Elder,
    &publisher.MinisterialServant, &publisher.RegularPionner)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
