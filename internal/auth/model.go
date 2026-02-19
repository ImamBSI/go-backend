package auth

type Account struct {
	ID   uint   `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	Role string `gorm:"column:role"`
}

func (Account) TableName() string {
	return "Account"
}

type User struct {
	ID        uint    `gorm:"column:id;primaryKey"`
	Username  string  `gorm:"column:username;unique;not null"`
	Password  string  `gorm:"column:password;not null"`
	AccountID uint    `gorm:"column:account_id"`
	Account   Account `gorm:"foreignKey:AccountID;references:ID"`
}

func (User) TableName() string {
	return "User"
}
