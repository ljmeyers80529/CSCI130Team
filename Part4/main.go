package finalWeb

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var cookieName string = "finalCookie"
var urlUserID = "userId"

const gcsBucket = "csci130-1265.appspot.com"

func init() {
	configureResourceLocation("css", "css")
	configureResourceLocation("img", "images")
	configureResourceLocation("ajax", "ajax")
	http.Handle("/favicon.ico", http.NotFoundHandler())    // ignore favico re quest (error 404)
	http.HandleFunc("/", index)                            // handle main page.
	http.HandleFunc("/logout", userLogout)                 // user log out.
	http.HandleFunc("/login", userLogin)                   // handle user login page.
	http.HandleFunc("/about", about)                       // about web page.
	http.HandleFunc("/register", register)                 // register new user.
	http.HandleFunc("/username/check", usernameCheck)      // verify username is unique.
	http.HandleFunc("/update", userDataUpdate)             // user informatuion update.
	http.HandleFunc("/upload", userFileUpload)             // upload file (txt, pdf, jpg, jpeg) to gcs.
	http.HandleFunc("/download", userFileDownload)         // download file from gcs.
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
