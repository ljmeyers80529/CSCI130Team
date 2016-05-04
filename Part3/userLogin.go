package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
	"strings"
)

func userLogin(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)
	userId, err := getID(res, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index getID: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Method == "POST" { 
	//&& req.FormValue("password") == "csci130" {

		var userFound bool = false // get user information from datastore.
		var ui userInformation

		q := datastore.NewQuery("Photos")
		t := q.Run(ctx)
		for {
			key, err := t.Next(&ui)
			if err == datastore.Done {
				break
			}
			if strings.ToLower(ui.Username) == strings.ToLower(req.FormValue("username")) {
				userFound = true
				userId = strings.Split(key.String(),",")[1]
				log.Infof(ctx, " .........................................................Key %v", key)
				log.Infof(ctx, " .........................................................Key1 %v", userId)
				log.Infof(ctx, " .........................................................Key2 %v", key.Kind())
				break
			}
		}
			log.Infof(ctx, "User found? : %v information = %v", userFound, ui)

		if userFound && ui.Password == req.FormValue("password") {
			if userId != "" {
				ui, err := retrieveUserInformationMemcache(ctx, userId, req)
				if err != nil {
					log.Errorf(ctx, "ERROR index retrieveMemc: %s", err) // expired cookie may exist on client
					http.Redirect(res, req, "/logout", http.StatusSeeOther)
					return
				}
				ui.LoggedIn = true
				ui.Username = req.FormValue("username")
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
		}
	}
	tpl.ExecuteTemplate(res, "login.html", nil)
}
