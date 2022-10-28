package models

import "github.com/kenfav/cong-reports-go/db"

func Update(id int64, publisher Publisher) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE publishers SET name=$2, birthdate=$3, baptismdate=$4, othersheep=$5, regularpionner=$6, elder=$7, ministerialservant=$8, surname=$9 WHERE id=$1`, 
		id, &publisher.Name, &publisher.BirthDate, &publisher.BaptismDate, &publisher.OtherSheep, &publisher.RegularPionner, &publisher.Elder, &publisher.MinisterialServant,  &publisher.Surname)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
