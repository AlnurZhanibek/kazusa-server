package main

import (
	"fmt"
	_ "github.com/AlnurZhanibek/kazusa-server/docs"
	"github.com/AlnurZhanibek/kazusa-server/internal/database"
	"github.com/AlnurZhanibek/kazusa-server/internal/handler"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/AlnurZhanibek/kazusa-server/internal/server"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
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

	moduleRepo := repository.NewModuleRepo(db)
	moduleService := service.NewModuleService(moduleRepo)
	moduleHandler := handler.NewModuleHandler(moduleService)

	courseRepo := repository.NewCourseRepo(db)
	courseService := service.NewCourseService(courseRepo, moduleRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	userRepo := repository.NewUserRepo(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	server.Start(&server.Handlers{
		CourseHandler: courseHandler,
		ModuleHandler: moduleHandler,
		AuthHandler:   authHandler,
	})

	fmt.Println("asshole 2")
}
