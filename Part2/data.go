package finalWeb

import (
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

type userInformation struct {
	UserId   string
	Name     string
	LoggedIn bool
}

// get an UUID from user.
func generateUUID() (string, error) {
	uuid, err := uuid.NewV4()
	return uuid.String(), err
}

func newUser(req *http.Request) (*http.Cookie, error) {
	id, err := generateUUID()
	if err != nil {
		ctx := appengine.NewContext(req)
		log.Errorf(ctx, "ERROR newVisitor uuid.NewV4: %s", err)
		return nil, err
	}
	m := initializeUserData(id)
	return makeCookie(m, req)
}

// set new user's information to default values
func initializeUserData(id string) userInformation {
	ui := userInformation{
		UserId:   id,
		LoggedIn: false,
		Name:     "",
	}
	return ui
}
