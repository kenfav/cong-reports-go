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
