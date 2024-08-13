package service

import (
	"fmt"
	"github.com/google/uuid"
	"kazusa-server/internal/entity"
	"kazusa-server/internal/repository"
)

type CourseServiceImplementation interface {
	Create(course entity.NewCourse) (bool, error)
	Read(pagination entity.Pagination, filters entity.CourseFilters) ([]entity.Course, error)
}

type CourseService struct {
	repo       repository.CourseRepositoryImplementation
	moduleRepo repository.ModuleRepositoryImplementation
}

func NewCourseService(repo repository.CourseRepositoryImplementation, moduleRepo repository.ModuleRepositoryImplementation) *CourseService {
	return &CourseService{repo: repo, moduleRepo: moduleRepo}
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

	if filters.ID != uuid.Nil || len(courses) == 1 {
		modules, err := s.moduleRepo.Read(entity.ModuleFilters{
			CourseID: courses[0].ID,
		}, entity.Pagination{})

		if err != nil {
			return nil, fmt.Errorf("course service get modules error: %v", err)
		}

		courses[0].Modules = &modules
	}

	return courses, nil
}
