package model

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username" gorm:"type:string; size:20; unique; not null"`
	Password string `json:"password" gorm:"type:string"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
