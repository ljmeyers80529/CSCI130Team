package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

type userFiles struct {
	Error    string
	UserFile []string
}

func userFileUpload(res http.ResponseWriter, req *http.Request) {
	var feedback userFiles

	ctx := appengine.NewContext(req)
	feedback.Error = ""

	userId, err := getID(res, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index getID: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	ui, err := retrieveUserInformationMemcache(ctx, userId, req)
	feedback.UserFile = ui.Files
	if err != nil {
		log.Errorf(ctx, "ERROR getting user information: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if req.Method == "POST" {
		feedback.Error = uploadFile(ctx, res, req, &ui)
		if feedback.Error == "" {
			feedback.UserFile = ui.Files
		}
	}

	if req.FormValue("id") == "" {
		http.Redirect(res, req, `/upload?id=`+userId, http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(res, "upload.html", feedback)
}
