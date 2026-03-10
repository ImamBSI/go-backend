package register

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
