package service

import (
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/google/uuid"
)

type CourseServiceImplementation interface {
	Create(course entity.NewCourse) (bool, error)
	Read(pagination entity.Pagination, filters entity.CourseFilters) ([]entity.Course, error)
	Update(body entity.CourseUpdateBody) (bool, error)
	Delete(id uuid.UUID) (bool, error)
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

func (s *CourseService) Update(course entity.CourseUpdateBody) (bool, error) {
	ok, err := s.repo.Update(course)
	if err != nil {
		return false, fmt.Errorf("course service update error: %v", err)
	}

	return ok, nil
}

func (s *CourseService) Delete(id uuid.UUID) (bool, error) {
	ok, err := s.repo.Delete(id)
	if err != nil {
		return false, fmt.Errorf("course service delete error: %v", err)
	}

	return ok, nil
}
