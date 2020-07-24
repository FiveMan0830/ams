package account

type AccountManagement interface {
	AddUser (username, surname, password string)
	AddGroup () 
}