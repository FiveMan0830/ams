package model

type RoleRelation struct {
	TeamID string `gorm:"primaryKey;column:team_id"`
	UserID string `gorm:"primaryKey;column:user_id"`
	Role   int    `gorm:"column:role;type:int(11)"`
}

func (RoleRelation) TableName() string {
	return "role_relation"
}
