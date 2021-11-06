package repository

// import (
// 	"testing"

// 	"github.com/go-playground/assert/v2"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// 	"ssl-gitlab.csie.ntut.edu.tw/ois/ois-project/ams/internal/model"
// )

// func setupMemoryDB() (*gorm.DB, TeamRepository) {
// 	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
// 	db.AutoMigrate(model.User{})
// 	db.AutoMigrate(model.Team{})
// 	db.AutoMigrate(model.RoleRelation{})
// 	teamRepo := NewTeamRepository(db)

// 	return db, teamRepo
// }

// func newDummyUser(id string) *model.User {
// 	return &model.User{
// 		ID:          "dummy_id_" + id,
// 		Account:     "dummy_account_" + id,
// 		DisplayName: "dummy_display_name",
// 		Password:    "dummy_password",
// 		Email:       "dummy_email",
// 	}
// }

// func newDummyTeam(id string) *model.Team {
// 	return &model.Team{
// 		ID:   id,
// 		Name: "dummy_team_name",
// 	}
// }

// func TestAddMembers(t *testing.T) {
// 	db, repo := setupMemoryDB()

// 	// prepare data
// 	teamId := "dummy_team"
// 	team := newDummyTeam(teamId)

// 	users := []*model.User{}
// 	users = append(users, newDummyUser("1"))
// 	users = append(users, newDummyUser("2"))
// 	users = append(users, newDummyUser("3"))

// 	// insert team and user into db
// 	db.Create(team)
// 	for _, user := range users {
// 		db.Create(user)
// 	}

// 	// make RoleRelation object
// 	roles := []int{1, 0, 3}
// 	relations := []model.RoleRelation{}
// 	for i, user := range users {
// 		relations = append(relations, model.RoleRelation{
// 			TeamID: team.ID,
// 			UserID: user.ID,
// 			Role:   roles[i],
// 		})
// 	}

// 	// add members into db
// 	if err := repo.AddMembers(relations); err != nil {
// 		t.Errorf("failed to add members")
// 	}

// 	// get team and make assertion
// 	var result model.Team
// 	db.Preload("Members").Where("id = ?", teamId).Find(&result)

// 	assert.Equal(t, "dummy_id_1", result.Members[0].User.ID)
// 	assert.Equal(t, "dummy_id_2", result.Members[1].User.ID)
// 	assert.Equal(t, "dummy_id_3", result.Members[2].User.ID)
// }
