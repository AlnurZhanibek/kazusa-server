package main

import (
	"fmt"
	"kazusa-server/internal/database"
	"kazusa-server/internal/entity"
	"kazusa-server/internal/handler"
	"kazusa-server/internal/repository"
	"kazusa-server/internal/server"
	"kazusa-server/internal/service"
	"log"
)

func main() {
	db := database.New()
	defer db.Close()

	courseRepo := repository.NewCourseRepo(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	server.Start(&server.Handlers{
		CourseHandler: courseHandler,
	})

	result, err := courseService.Read(entity.Pagination{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		log.Fatalf("some error: %v", err)
	}

	fmt.Printf("%v", result)
}
