package models

type User struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Subscribe bool   `json:"subscribe"`
	Points    int    `json:"points"`
}
