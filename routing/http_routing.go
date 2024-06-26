package routing

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path"
	"text/template"
)

const BASE_PAGE_PATH = "templates/pages/"
const BASE_PARTS_PATH = "templates/parts/"

const BASE_HTML = BASE_PARTS_PATH + "base.html"
const HEADER_HTML = BASE_PARTS_PATH + "header.html"
const FOOTER_HTML = BASE_PARTS_PATH + "footer.html"

const POST_HTML = BASE_PAGE_PATH + "post.html"
const POSTLIST_HTML = BASE_PAGE_PATH + "postlist.html"
const ERROR_HTML = BASE_PAGE_PATH + "error.html"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// all paths fall back to "/", that's why we handle all 404 status codes like this
	if r.URL.Path != "/" {
		r.Header.Set("Status", fmt.Sprint(http.StatusNotFound))
		ErrorHandler(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles(BASE_HTML, POST_HTML, HEADER_HTML, FOOTER_HTML))

	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func PostlistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(BASE_HTML, POSTLIST_HTML, HEADER_HTML, FOOTER_HTML))

	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(BASE_HTML, ERROR_HTML, HEADER_HTML, FOOTER_HTML))

	statusHeader := r.Header.Get("Status")
	fmt.Println("Error Code: ", statusHeader)

	err := tmpl.ExecuteTemplate(w, "base", statusHeader)
	if err != nil {
		fmt.Println(err)
	}
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	suffix := path.Ext(url)
	mimeType := mime.TypeByExtension(suffix)
	w.Header().Set("Content-Type", mimeType)
	w.Header().Set("Cache-Control", "max-age=10800")

	data, err := os.ReadFile(url[1:])
	if err != nil {
		fmt.Print(err)
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Print(err)
	}
}
