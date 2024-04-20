package server

import (
	"TownVoice/internal/handlers"
	"html/template"
	"log"
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
			Body:  "This is dynamic content from the server.go",
		}

		tmpl.Execute(w, data)
	})
	http.HandleFunc("/login", handlers.LoginHandler)

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
	println("Server started on :8082")
}
