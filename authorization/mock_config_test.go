package authorization

type mockConfig struct {
	tokenSecret string
	expiredDays int
}