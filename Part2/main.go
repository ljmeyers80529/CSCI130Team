package finalWeb

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var cookieName string = "finalCookie"
var urlUserID = "userId"

// map resource physical location ro href relative location.
// phyDir : resource files physical location relative to html file.
// hrefDir: resource location as defined withing the href tag.
func configureResourceLocation(phyDir, hrefDir string) {
	fs := http.FileServer(http.Dir(phyDir))
	fs = http.StripPrefix("/"+hrefDir, fs)
	http.Handle("/"+hrefDir+"/", fs)
}

// render main web page if someone has logged in.
func index(res http.ResponseWriter, req *http.Request) {
	userId := getID(res, req)
	user := retrieveUserData(userId, req)
	if user.LoggedIn {
		tpl.ExecuteTemplate(res, "index.html", user) // executer main page HTML template
	} else {
		http.Redirect(res, req, "/login?id="+userId, http.StatusFound)
	}
}

func init() {
	configureResourceLocation("css", "css")
	configureResourceLocation("img", "images")
	http.Handle("/favicon.ico", http.NotFoundHandler())    // ignore favico re quest (error 404)
	http.HandleFunc("/", index)                            // handle main page.
	http.HandleFunc("/login", userLogin)                   // handle user login page.
	tpl = template.Must(template.ParseGlob("html/*.html")) // load and parse all web pages for this project.
}
