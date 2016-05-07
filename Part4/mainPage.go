package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

// render main web page if someone has logged in.
func index(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	userId, err := getID(res, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index getID: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// 	if req.Method == "POST" {                   //todo
	// 	}

	ui, err := retrieveUserInformationMemcache(ctx, userId, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index retrieveMemc: %s", err) // may have an expired cookie.
		http.Redirect(res, req, "/logout", http.StatusSeeOther)
		return
	}
	if ui.LoggedIn {
		tpl.ExecuteTemplate(res, "index.html", ui)
	} else {
		userLogin(res, req)
	}

}
