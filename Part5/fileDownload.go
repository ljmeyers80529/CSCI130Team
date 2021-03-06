package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

type dFile struct {
	Name string
	Link string
	Type string
}

func userFileDownload(res http.ResponseWriter, req *http.Request) {
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

	f := fileList(ctx, res, req, &ui)

	// if req.Method == "POST" && user != "" && req.FormValue("password") == req.FormValue("confirm") {
	// }

	if req.FormValue("id") == "" {
		http.Redirect(res, req, `/download?id=`+userId, http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(res, "download.html", f)
}
