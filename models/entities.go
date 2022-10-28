package models

type Publisher struct {
	ID          int64  `json:"id"`
  Surname string `json:"surname"`
	Name        string `json:"name"`
	BirthDate   int64  `json:"birthdate"`
	BaptismDate int64  `json:"baptismdate"`
	OtherSheep  bool   `json:"othersheep"`
  Elder bool `json:"elder"`
  MinisterialServant bool `json:"ministerialservant"`
  RegularPionner bool `json:"regularpionner"`
}
