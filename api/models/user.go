package models

type User struct {
	IDUser       uint64 `gorm:"column:id_user;primaryKey;autoIncrement"`
	UserEmail    string `gorm:"column:user_email;size:160"`
	UserPassword string `gorm:"column:user_password;size:255"`
	UserStatus   string `gorm:"column:user_status;size:4"`
	UserMerchant bool   `gorm:"column:user_merchant"`
}

// TableName overrides the default table name
func (User) TableName() string {
	return "users"
}
