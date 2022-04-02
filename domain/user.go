package domain

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string `gorm:"name" json:"name"`
	Contact  string `gorm:"contact" json:"contact"`
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
	Address  string `gorm:"address" json:"address"`
}

type UserRepository interface {
	Register(*User) (*User, error)
	Delete(*User) error
	Update(*User, *User) (*User, error)
}
