package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"`
}
