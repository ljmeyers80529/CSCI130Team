package finalWeb

type UserInformation struct {
	UserId   string
	Name     string
	LoggedIn bool
}

// set new user's information to default values
func initializeUserData(id string) UserInformation {
	ui := UserInformation{
		UserId:   id,
		LoggedIn: false,
		Name:     "",
	}
	return ui
}
