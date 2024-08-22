package service

import (
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/google/uuid"
)

type ModuleServiceImplementation interface {
	Create(Module entity.NewModule) (bool, error)
	Read(pagination entity.Pagination, filters entity.ModuleFilters) ([]entity.Module, error)
	Update(body entity.ModuleUpdateBody) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}

type ModuleService struct {
	repo repository.ModuleRepositoryImplementation
}

func NewModuleService(repo repository.ModuleRepositoryImplementation) *ModuleService {
	return &ModuleService{repo: repo}
}

func (s *ModuleService) Create(Module entity.NewModule) (bool, error) {
	ok, err := s.repo.Create(Module)
	if err != nil {
		return false, fmt.Errorf("module service create error: %v", err)
	}

	return ok, nil
}

func (s *ModuleService) Read(pagination entity.Pagination, filters entity.ModuleFilters) ([]entity.Module, error) {
	modules, err := s.repo.Read(filters, pagination)
	if err != nil {
		return nil, fmt.Errorf("module service create error: %v", err)
	}

	return modules, nil
}

func (s *ModuleService) Update(body entity.ModuleUpdateBody) (bool, error) {
	ok, err := s.repo.Update(body)
	if err != nil {
		return false, fmt.Errorf("module service update error: %v", err)
	}

	return ok, nil
}

func (s *ModuleService) Delete(id uuid.UUID) (bool, error) {
	ok, err := s.repo.Delete(id)
	if err != nil {
		return false, fmt.Errorf("module service delete error: %v", err)
	}

	return ok, nil
}
