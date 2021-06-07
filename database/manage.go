package database

// import "database/sql"

type Management interface {
	GetRole(userID, teamID string) (int, error) 
}