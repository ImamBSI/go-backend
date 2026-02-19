package register

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
	Role string `gorm:"column:role"`
	// Accounts []Account `gorm:"foreignKey:UserID"`
}

func (User) TableName() string { return "User" }

type Account struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"column:username;unique;not null"`
	Password string `gorm:"column:password;not null"`
	UserID   uint   `gorm:"column:user_id"`
	User     User   `gorm:"foreignKey:UserID;references:ID"`
}

func (Account) TableName() string { return "Account" }
