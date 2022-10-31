package models

type Publisher struct {
	ID          int64  `json:"id"`
  Surname string `json:"surname"`
	Name        string `json:"name"`
	BirthDate   int64  `json:"birthdate"`
	BaptismDate int64  `json:"baptismdate"`
  OtherSheep  bool   `json:"othersheep" default:"true"`
  Elder bool `json:"elder" default:"false"`
  MinisterialServant bool `json:"ministerialservant" default:"false"`
  RegularPionner bool `json:"regularpionner" default:"false"`
}

type Report struct {
  ID int64 `json:"id"`
  PublisherId int64 `json:"publishers_id"`
  Year int64 `json:"date_year"`
  Month int64 `json:"date_month"`
  Publications int64 `json:"publications"`
  Videos int64 `json:"videos"`
  Hours int64 `json:"hours"`
  ReturnVisits int64 `json:"return_visits"`
  BibleStudies int64 `json:"bible_studies"`
  Notes string `json:"notes"`
}
