package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func makeCookie(ui userInformation, req *http.Request) (*http.Cookie, error) {
	ctx := appengine.NewContext(req)

		// DATASTORE
		err := setUserInformationDatastore(ctx, ui, req)
		if err != nil {
			log.Errorf(ctx, "ERROR makeCookie storeDstore: %s", err)
			return nil, err
		}

	// MEMCACHE
	err = setUserInformationMemcache(ctx, ui, req)
	if err != nil {
		log.Errorf(ctx, "ERROR makeCookie storeMemc: %s", err)
		return nil, err
	}

	// COOKIE
	cookie := &http.Cookie{
		Name:  cookieName,
		Value: ui.UserId,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie, nil
}

// get current user loggin in.
func currentUser(ui userInformation, req *http.Request) (*http.Cookie, error) {
	return makeCookie(ui, req)
}
