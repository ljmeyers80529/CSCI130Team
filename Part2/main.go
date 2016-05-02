package finalWeb

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var cookieName string = "finalCookie"
var urlUserID = "userId"

func init() {
	configureResourceLocation("css", "css")
	configureResourceLocation("img", "images")
	http.Handle("/favicon.ico", http.NotFoundHandler())    // ignore favico re quest (error 404)
	http.HandleFunc("/", index)                            // handle main page.
	http.HandleFunc("/logout", userLogout)                 // user log out.
	http.HandleFunc("/login", userLogin)                   // handle user login page.
        http.HandleFunc("/about", about)                       // about web page.
	tpl = template.Must(template.ParseGlob("html/*.html")) // load and parse all web pages for this project.
}

// map resource physical location ro href relative location.
// phyDir : resource files physical location relative to html file.
// hrefDir: resource location as defined withing the href tag.
func configureResourceLocation(phyDir, hrefDir string) {
	fs := http.FileServer(http.Dir(phyDir))
	fs = http.StripPrefix("/"+hrefDir, fs)
	http.Handle("/"+hrefDir+"/", fs)
}
