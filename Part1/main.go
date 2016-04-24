package finalWeb

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	configureResourceLocation("css", "css")
	configureResourceLocation("images", "img")
	http.Handle("/favicon.ico", http.NotFoundHandler())      // ignore favico re quest (error 404)
	http.HandleFunc("/", index)                              // handle main page.
	tpl = template.Must(template.ParseGlob("html/*.gohtml")) // load and parse all web pages for this project.
}

// map resource physical location ro href relative location.
// phyDir : resource files physical location relative to html file.
// hrefDir: resource location as defined withing the href tag.
func configureResourceLocation(phyDir, hrefDir string) {
	fs := http.FileServer(http.Dir(phyDir))
	fs = http.StripPrefix("/"+hrefDir, fs)
	http.Handle("/"+hrefDir+"/", fs)
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil) // executer main page HTML template
}
