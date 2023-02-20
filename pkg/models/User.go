package models

type Point struct {
	Id     int `json:"id" gorm:"primaryKey"`
	Points int `json:"points"`
}
