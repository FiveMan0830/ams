package database

// import "database/sql"

type Management interface {
	InsertRole(userID string, teamID string, role int)
	GetRole(userID, teamID string) (int, error) 
}