package models

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
