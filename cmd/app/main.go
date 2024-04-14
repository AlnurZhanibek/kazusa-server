package main

import (
	"kazusa-server/internal/database"
	"kazusa-server/internal/handler"
	"kazusa-server/internal/repository"
	"kazusa-server/internal/server"
	"kazusa-server/internal/service"
)

func main() {
	db := database.New()
	defer db.Close()

	courseRepo := repository.NewCourseRepo(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	moduleRepo := repository.NewModuleRepo(db)
	moduleService := service.NewModuleService(moduleRepo)
	moduleHandler := handler.NewModuleHandler(moduleService)

	server.Start(&server.Handlers{
		CourseHandler: courseHandler,
		ModuleHandler: moduleHandler,
	})
}
