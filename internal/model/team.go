package model

type Team struct {
	ID      string  `gorm:"primaryKey;column:id;type:varchar(36)"`
	Name    string  `gorm:"type:varchar(255)"`
	Members []*User `gorm:"many2many:role_relation"`
}

func (Team) TableName() string {
	return "team"
}
