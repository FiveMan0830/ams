package model

type Member struct {
	User User  `json:"user"`
	Role int64 `json:"role"`
}

type Team struct {
	ID       string    `gorm:"primaryKey;column:id;type:varchar(36)"`
	Name     string    `gorm:"type:varchar(255)"`
	Members  []*Member `gorm:"-"`
	Subteams []*Team   `gorm:"many2many:team_relation" json:"-"`
}

func (Team) TableName() string {
	return "team"
}
