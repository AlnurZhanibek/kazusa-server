package server

import (
	"fmt"
	"github.com/rs/cors"
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
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("/course", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.CourseHandler.Read(w, r)
		case http.MethodPost:
			handlers.CourseHandler.Create(w, r)
		}
	})

	mux.HandleFunc("/module", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ModuleHandler.Read(w, r)
		case http.MethodPost:
			handlers.ModuleHandler.Create(w, r)
		}
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.AuthHandler.Login(w, r)
		}
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.AuthHandler.Register(w, r)
		}
	})

	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.AuthHandler.Logout(w, r)
		}
	})

	mux.HandleFunc("/swagger", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))

	muxHandler := cors.Default().Handler(mux)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), muxHandler)
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
