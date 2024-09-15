package main

import (
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

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	activityRepo := repository.NewActivityRepository(db)
	activityService := service.NewActivityService(activityRepo)
	activityHandler := handler.NewActivityHandler(activityService)

	moduleRepo := repository.NewModuleRepo(db)
	moduleService := service.NewModuleService(moduleRepo, activityService)
	moduleHandler := handler.NewModuleHandler(moduleService)

	fileService := service.NewFileService()

	courseRepo := repository.NewCourseRepo(db)
	courseService := service.NewCourseService(courseRepo, moduleService, fileService)
	courseHandler := handler.NewCourseHandler(courseService)

	server.Start(&server.Handlers{
		CourseHandler:   courseHandler,
		ModuleHandler:   moduleHandler,
		UserHandler:     userHandler,
		AuthHandler:     authHandler,
		ActivityHandler: activityHandler,
	})
}
