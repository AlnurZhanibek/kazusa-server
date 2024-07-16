package server

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"kazusa-server/internal/handler"
	"log"
	"net/http"
	"os"
)

type Handlers struct {
	CourseHandler *handler.CourseHandler
	ModuleHandler *handler.ModuleHandler
	AuthHandler   *handler.AuthHandler
}

func Start(handlers *Handlers) {
	port := os.Getenv("HTTP_PORT")

	http.HandleFunc("/course", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.CourseHandler.Read(w, r)
		case http.MethodPost:
			handlers.CourseHandler.Create(w, r)
		}
	})

	http.HandleFunc("/module", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ModuleHandler.Read(w, r)
		case http.MethodPost:
			handlers.ModuleHandler.Create(w, r)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.AuthHandler.Login(w, r)
		}
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.AuthHandler.Register(w, r)
		}
	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.AuthHandler.Logout(w, r)
		}
	})

	http.HandleFunc("/swagger", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
