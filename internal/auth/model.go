package auth

type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Role string `gorm:"column:role" json:"role"`
}

func (User) TableName() string { return "User" }

type Account struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"column:username;unique;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"-"`
	UserID   uint   `gorm:"column:user_id" json:"user_id"`
	User     User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (Account) TableName() string { return "Account" }
