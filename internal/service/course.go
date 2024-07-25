package service

import (
	"fmt"
	"kazusa-server/internal/entity"
	"kazusa-server/internal/repository"
)

type CourseServiceImplementation interface {
	Create(course entity.NewCourse) (bool, error)
	Read(pagination entity.Pagination, filters entity.CourseFilters) ([]entity.Course, error)
}

type CourseService struct {
	repo repository.CourseRepositoryImplementation
}

func NewCourseService(repo repository.CourseRepositoryImplementation) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Create(course entity.NewCourse) (bool, error) {
	ok, err := s.repo.Create(course)
	if err != nil {
		return false, fmt.Errorf("course service create error: %v", err)
	}

	return ok, nil
}

func (s *CourseService) Read(pagination entity.Pagination, filters entity.CourseFilters) ([]entity.Course, error) {
	courses, err := s.repo.Read(pagination, filters)
	if err != nil {
		return nil, fmt.Errorf("course service create error: %v", err)
	}

	return courses, nil
}
