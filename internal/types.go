package internal

type Member struct {
	ID string `json:"id"`
	Account string `json:"account"`
	DisplayName string `json:"displayName"`
	Email string `json:"email"`
	Role string `json:"role"`
}