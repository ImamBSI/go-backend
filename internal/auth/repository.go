package auth

import (
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) FindAccountByUsername(username string) (*Account, error) {
	var account Account
	if err := r.Db.Preload("User").Where("username = ?", username).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *Repository) IsUsernameExists(username string) bool {
	var account Account
	if err := r.Db.Where("username = ?", username).First(&account).Error; err == nil {
		return true
	}
	return false
}

func (r *Repository) CreateUser(user *User) error {
	return r.Db.Create(user).Error
}

func (r *Repository) CreateAccount(account *Account) error {
	return r.Db.Create(account).Error
}

func (r *Repository) GetUsers(page int, limit int) ([]User, int64, error) {
	var users []User
	var total int64

	offset := (page - 1) * limit

	if err := r.Db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.Db.
		Limit(limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *Repository) GetUserByID(id uint) (*User, error) {

	var user User

	if err := r.Db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) DeleteUser(id uint) error {

	err := r.Db.Delete(&User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
