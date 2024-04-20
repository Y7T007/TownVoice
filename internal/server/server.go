package server

import (
	"TownVoice/internal/auth/controller"
	"TownVoice/internal/handlers"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
	Body  string
}

func Start(port string) {
	println("Server started on :" + port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./web/index.html"))

		data := PageData{
			Title: "My Dynamic Title",
			Body:  "This is dynamic content from the server.go",
		}

		tmpl.Execute(w, data)
	})
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/auth/register-client", controller.RegisterClient)
	http.HandleFunc("/auth/login-client", controller.LoginClient)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
	println("Server started on :" + port)
}
