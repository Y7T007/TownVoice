package server

import (
	"TownVoice/internal/auth/controller"
	"TownVoice/internal/handlers"
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
	Body  string
}

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./web/index.html"))

		data := PageData{
			Title: "My Dynamic Title",
			Body:  "This is dynamic content from the server.go",
		}

		tmpl.Execute(w, data)
	})
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/auth/register-client", controller.RegisterClient)
	mux.HandleFunc("/auth/login-client", controller.LoginClient)

	return mux
}
