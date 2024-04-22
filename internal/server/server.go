package server

import (
	"TownVoice/internal/auth/controller"
	"TownVoice/internal/handlers"
	"TownVoice/internal/routes"
	"TownVoice/utils"

	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Title string
	Body  string
}

func getNumberHandler(w http.ResponseWriter, r *http.Request) {
	numbers := make(map[string][]int)
	for i := 1; i <= 10; i++ {
		numbers["numbers"] = append(numbers["numbers"], i)
	}

	json.NewEncoder(w).Encode(numbers)
}
func gettoken(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		http.Error(w, "Authorization header not provided", http.StatusUnauthorized)
		return
	}

	idToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

	token, err := utils.VerifyIDToken(r.Context(), idToken)
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}

	// You can access the user's Firebase UID as follows:
	uid := token.UID
	fmt.Fprintf(w, "User with UID %s authenticated\n", uid)
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
	mux.Handle("/get_number", utils.Middleware(http.HandlerFunc(getNumberHandler)))
	mux.Handle("/verifytoken", utils.Middleware(http.HandlerFunc(gettoken)))

	routes.CommentRoutes(mux)

	return mux
}
