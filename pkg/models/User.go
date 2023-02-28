package models

type Point struct {
	IdUser int `json:"id" gorm:"primaryKey;column:IdUser" `
	Points int `json:"points gorm:"column:Points"`
}
