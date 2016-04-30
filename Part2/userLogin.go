package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func userLogin(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" && req.FormValue("password") == "csci130" {
		var ui UserInformation
		ui.Name = req.FormValue("username")
		ui.UserId = getID(res, req)
		ui.LoggedIn = true
		setUserData(ui, req)
		http.Redirect(res, req, "/", http.StatusFound)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", nil) // executer main page HTML template
}
