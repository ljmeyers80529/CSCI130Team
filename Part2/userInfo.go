package finalWeb

import (
	"github.com/nu7hatch/gouuid"
	// "google.golang.org/appengine"
	// "google.golang.org/appengine/log"
	// "google.golang.org/appengine/memcache"
	"net/http"
)

// get an UUID from user.
func generateUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

// get user id either from a cookie or the URL.
func getID(res http.ResponseWriter, req *http.Request) string {
	var cookie *http.Cookie // new cookie.
	var userId string

	cookie, err := req.Cookie(cookieName) // get if there is a cookie already present.
	if err != nil {
		userId = req.FormValue(userId) // no cookie, check URL string.
		if userId == "" {
			userId = generateUUID()
		}
		cookie = &http.Cookie{ // try to store id for later use in COOKIE
			Name:  cookieName,
			Value: userId,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	userId = cookie.Value
	return userId
}

func retrieveUserData(userId string, req *http.Request) (userInfo UserInformation) {
	ui, _ := retrieveUserInformationMemcache(userId, req)
	return ui
}

func setUserData(userInfo UserInformation, req *http.Request) {
	setUserInformationMemcache(userInfo, req)
}
