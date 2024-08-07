package main

import (
	_ "kazusa-server/docs"
	"kazusa-server/internal/database"
	"kazusa-server/internal/handler"
	"kazusa-server/internal/repository"
	"kazusa-server/internal/server"
	"kazusa-server/internal/service"
)

// @title Swagger KazUSA API
// @version 1.0
// @description This is the KazUSA server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://kazusa.kz
// @contact.email aln.zh.621@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host kazusa.kz
// @BasePath /
func main() {
	db := database.New()
	defer db.Close()

	courseRepo := repository.NewCourseRepo(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	moduleRepo := repository.NewModuleRepo(db)
	moduleService := service.NewModuleService(moduleRepo)
	moduleHandler := handler.NewModuleHandler(moduleService)

	userRepo := repository.NewUserRepo(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	server.Start(&server.Handlers{
		CourseHandler: courseHandler,
		ModuleHandler: moduleHandler,
		AuthHandler:   authHandler,
	})
}
