package finalWeb

import (
	// 	"golang.org/x/net/context"
	// 	"google.golang.org/appengine/log"
	// 	"google.golang.org/appengine/datastore"
	"errors"
	"net/http"
	"net/mail"
)

// get user id either from a cookie or the URL.
func getID(res http.ResponseWriter, req *http.Request) (string, error) {

	var id string
	var cookie *http.Cookie

	cookie, err := req.Cookie(cookieName) // try to get the id from the COOKIE
	if err == http.ErrNoCookie {

		id := req.FormValue(urlUserID) // try to get the id from the URL
		if id == "" {
			http.Redirect(res, req, "/logout", http.StatusSeeOther) // no id, so create one BRAND NEW
			return id, errors.New("ERROR: redirect to /logout because no session id accessible")
		}
		cookie = &http.Cookie{ // try to store id for later use in COOKIE
			Name:  cookieName,
			Value: id,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	id = cookie.Value
	return id, nil
}

// check email address is valid.
func emailAddressValid(emailAddr string) bool {
	_, err := mail.ParseAddress(emailAddr)
	return err == nil
}
