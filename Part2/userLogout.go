package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func userLogout(res http.ResponseWriter, req *http.Request) {
	cookie, err := newUser(req)
	if err != nil {
		ctx := appengine.NewContext(req)
		log.Errorf(ctx, "ERROR logout getCookie: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(res, cookie)
	http.Redirect(res, req, `/?id=`+cookie.Value, http.StatusSeeOther)
}
