package user

import (
	"BCC_Intership_BuildUlang/domain"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) Create(userRegister *domain.User) (*domain.User, error) {
	if err := u.db.Create(&userRegister).Error; err != nil {
		return nil, err
	}
	return userRegister, nil
}

func (u *user) Delete(id string) error {
	var userDel *domain.User
	if err := u.db.Where("id = ?", id).First(&userDel).Error; err != nil {
		return err
	}
	if err := u.db.Delete(&userDel).Error; err != nil {
		return err
	}
	return nil
}

func (u *user) Update(id string, userUpdate *domain.User) (*domain.User, error) {
	var userOld *domain.User
	if err := u.db.Where("id = ?", id).First(&userOld).Error; err != nil {
		return nil, err
	}
	if err := u.db.Model(&userOld).Updates(&userUpdate).Error; err != nil {
		return nil, err
	}
	return userUpdate, nil
}
