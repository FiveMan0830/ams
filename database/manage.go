package database

// import "database/sql"

type Management interface {
	InsertRole(userID string, teamID string, role int)
	GetRole(userID, teamID string) (int, error) 
	GetTeamLeader(teamID string) (string, error)
	UpdateLeader(oldLeaderID, newLeaderID, teamID string)
	DeleteTeam( teamID string)
	DeleteRole(userID, teamID string)
}