package server

import (
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/handler"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

type Handlers struct {
	CourseHandler *handler.CourseHandler
	ModuleHandler *handler.ModuleHandler
	UserHandler   *handler.UserHandler
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
		case http.MethodPut:
			handlers.CourseHandler.Update(w, r)
		case http.MethodDelete:
			handlers.CourseHandler.Delete(w, r)
		}
	})

	mux.HandleFunc("/module", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ModuleHandler.Read(w, r)
		case http.MethodPost:
			handlers.ModuleHandler.Create(w, r)
		case http.MethodPut:
			handlers.ModuleHandler.Update(w, r)
		case http.MethodDelete:
			handlers.ModuleHandler.Delete(w, r)
		}
	})

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.UserHandler.Read(w, r)
		case http.MethodPost:
			handlers.UserHandler.Create(w, r)
		case http.MethodPut:
			handlers.UserHandler.Update(w, r)
		case http.MethodDelete:
			handlers.UserHandler.Delete(w, r)
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

	mux.HandleFunc("/swagger", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))

	muxHandler := cors.Default().Handler(mux)

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), muxHandler)
	if err != nil {
		log.Fatalf("server error: %v", err)
	} else {
		log.Printf("server running on port: %v\n", port)
	}
}
