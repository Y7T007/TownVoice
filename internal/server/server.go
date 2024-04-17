package server

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
	Body  string
}

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./web/index.html"))

		data := PageData{
			Title: "My Dynamic Title",
			Body:  "This is dynamic content from the server.go ",
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
