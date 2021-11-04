package model

type User struct {
	ID          string `gorm:"primaryKey;column:id;type:varchar(36)"`
	Account     string `gorm:"uniqueIndex;column:account;type:varchar(255);not null"`
	DisplayName string `gorm:"type:varchar(255)"`
	Password    string `gorm:"column:password;type:varchar(50);not null" json:"-"`
	Email       string `gorm:"column:email;type:varchar(255)"`
}

func (User) TableName() string {
	return "user"
}
