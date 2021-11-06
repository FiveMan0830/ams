package model

type TeamRelation struct {
	TeamID    string `gorm:"primaryKey"`
	SubteamID string `gorm:"primaryKey"`
}

func (TeamRelation) TableName() string {
	return "team_relation"
}
