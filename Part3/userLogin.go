package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func userLogin(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)
	userId, err := getID(res, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index getID: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Method == "POST" && req.FormValue("password") == "csci130" {
		ui, err := retrieveUserInformationMemcache(ctx, userId, req)
		if err != nil {
			log.Errorf(ctx, "ERROR index retrieveMemc: %s", err) // expired cookie may exist on client
			http.Redirect(res, req, "/logout", http.StatusSeeOther)
			return
		}
		ui.LoggedIn = true
		ui.Name = req.FormValue("username")
		ui.UserId = userId

		cookie, err := currentUser(ui, req)
		if err != nil {
			log.Errorf(ctx, "ERROR login currentVisitor: %s", err)
			http.Redirect(res, req, `/?id=`+cookie.Value, http.StatusSeeOther)
			return
		}
		http.SetCookie(res, cookie)
		http.Redirect(res, req, `/?id=`+cookie.Value, http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", nil)
}
