package repository

import (
	"file-zipper-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByID(id uint) (*model.User, error)
	FindByGoogleSub(googleSub string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := ur.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) FindByGoogleSub(googleSub string) (*model.User, error) {
	var user model.User
	err := ur.db.Where("google_sub = ?", googleSub).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) Update(user *model.User) error {
	return ur.db.Save(user).Error
}
