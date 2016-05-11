package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func userDataUpdate(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)
	userId, err := getID(res, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index getID: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	ui, err := retrieveUserInformationMemcache(ctx, userId, req)
	if err != nil {
		log.Errorf(ctx, "ERROR getting user information: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	user := req.FormValue("username")
	if req.Method == "POST" && user != "" && req.FormValue("password") == req.FormValue("confirm") {
		commitNewUsername(res, req, user)

		if req.FormValue("email") == "" || emailAddressValid(req.FormValue("email")) {
			ui := userInformation{
				UserId:   userId,
				Username: user,
				Password: req.FormValue("password"),
				Name:     req.FormValue("name"),
				Email:    req.FormValue("email"),
				Age:      req.FormValue("age"),
				LoggedIn: true,
			}
			setUserInformationDatastore(ctx, ui, req)
			setUserInformationMemcache(ctx, ui, req)
			http.Redirect(res, req, "/", http.StatusSeeOther)
		} else {
			ui.Email = "<invalid>"
			setUserInformationDatastore(ctx, ui, req)
			setUserInformationMemcache(ctx, ui, req)
			http.Redirect(res, req, "/update", http.StatusSeeOther)
		}
	}
	if req.FormValue("id") == "" {
		http.Redirect(res, req, `/update?id=`+userId, http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(res, "update.html", ui)
}
