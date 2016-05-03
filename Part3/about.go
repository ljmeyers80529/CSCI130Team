package finalWeb

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func about(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)
	userId, err := getID(res, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index getID: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
        if req.FormValue("id") == "" {
		http.Redirect(res, req, `/about?id=`+userId, http.StatusSeeOther)
        }
	tpl.ExecuteTemplate(res, "about.html", nil)
}
