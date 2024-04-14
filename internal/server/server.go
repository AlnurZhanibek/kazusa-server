package server

import (
	"fmt"
	"kazusa-server/internal/handler"
	"log"
	"net/http"
	"os"
)

type Handlers struct {
	CourseHandler *handler.CourseHandler
	ModuleHandler *handler.ModuleHandler
}

func Start(handlers *Handlers) {
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

	err := http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("HTTP_PORT")), nil)
	if err != nil {
		log.Fatalf("server error: %v", err)
	}
}
