package routing

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path"
	"text/template"
)

const BasePagePath = "templates/pages/"
const BasePartsPath = "templates/parts/"

const BaseHTML = BasePartsPath + "base.html"
const HeaderHTML = BasePartsPath + "header.html"
const FooterHTML = BasePartsPath + "footer.html"

const PostHTML = BasePagePath + "post.html"
const PostlistHTML = BasePagePath + "postlist.html"
const GameHTML = BasePagePath + "game.html"
const ErrorHTML = BasePagePath + "error.html"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// all paths fall back to "/", that's why we handle all 404 status codes like this
	if r.URL.Path == "/game" {
		GameHandler(w, r)
		fmt.Println("Try to render game site from index handler...")
		return
	} else if r.URL.Path != "/" {
		r.Header.Set("Status", fmt.Sprint(http.StatusNotFound))
		fmt.Println(r.URL.Path)
		ErrorHandler(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles(BaseHTML, PostHTML, HeaderHTML, FooterHTML))
	fmt.Println("Rendering at: ", r.URL.Path)

	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(BaseHTML, GameHTML, HeaderHTML, FooterHTML))

	fmt.Println("Try to render game site...")

	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func PostlistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(BaseHTML, PostlistHTML, HeaderHTML, FooterHTML))

	fmt.Println("Error at Poslist/: ", r.URL.Path)

	err := tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(BaseHTML, ErrorHTML, HeaderHTML, FooterHTML))

	statusHeader := r.Header.Get("Status")
	fmt.Println("Error Code: ", statusHeader)
	fmt.Println("Error at: ", r.URL.Path)

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
