package routing

import (
	"fmt"
	"net/http"
	"text/template"
)

const BASE_PATH = "templates/pages/"
const BASE_HTML = "templates/parts/base.html"
const HEADER_HTML = "templates/parts/header.html"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// all paths fall back to "/", that's why we handle all 404 status codes like this
	if r.URL.Path != "/" {
		r.Header.Set("Status", fmt.Sprint(http.StatusNotFound))
		ErrorHandler(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles(BASE_HTML, HEADER_HTML, FRONTPAGE_HTML, FOOTER_HTML))

	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(BASE_HTML, HEADER_HTML, ERROR_HTML, FOOTER_HTML))

	header := r.Header.Get("Status")
	fmt.Println(header)

	err := tmpl.ExecuteTemplate(w, "base", header)
	if err != nil {
		fmt.Println(err)
	}
}
