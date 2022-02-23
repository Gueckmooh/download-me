package server

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var (
	templateFiles = []string{
		"html/home.html",
	}

	templates = template.Must(template.New("templates").ParseFiles(templateFiles...))
)

func renderTemplate(w io.Writer, tmpl string, data interface{}) error {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		return fmt.Errorf("renderTemplates: %s", err.Error())
	}
	return nil
}

func renderHome(wr io.Writer) error {
	data := struct {
		Downloadables []Downloadable
	}{
		Downloadables: getDownloadables(),
	}
	return renderTemplate(wr, "home", data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderHome(w); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
