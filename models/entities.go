package models

type Publisher struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	BirthDate   int64  `json:"birthdate"`
	BaptismDate int64  `json:"baptismdate"`
	OtherSheep  bool   `json:"othersheep"`
}
