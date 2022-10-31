package models

import (
	"github.com/kenfav/cong-reports-go/db"
)

func GetReportByMonth(publisher_id int64, year int64, month int64) (report Report, err error) {
  conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM reports WHERE publishers_id=$1 AND date_year=$2 AND date_month=$3`, publisher_id, year, month)

	err = row.Scan(
    &report.ID,
    &report.PublisherId,
    &report.Year,
    &report.Month,
    &report.Publications,
    &report.Videos,
    &report.Hours,
    &report.ReturnVisits,
    &report.BibleStudies,
    &report.Notes,
  )
	return
}
