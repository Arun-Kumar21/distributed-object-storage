package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"-"`
}
