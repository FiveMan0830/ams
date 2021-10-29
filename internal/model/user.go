package model

type User struct {
	ID          string `gorm:"primaryKey"`
	Account     string `gorm:"uniqueIndex"`
	DisplayName string 
	Password    string 
	Email       string 
}
