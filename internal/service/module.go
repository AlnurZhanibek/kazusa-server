package service

import (
	"fmt"
	"kazusa-server/internal/entity"
	"kazusa-server/internal/repository"
)

type ModuleServiceImplementation interface {
	Create(Module entity.NewModule) (bool, error)
	Read(filters entity.ModuleFilters) ([]entity.Module, error)
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

func (s *ModuleService) Read(filters entity.ModuleFilters) ([]entity.Module, error) {
	modules, err := s.repo.Read(filters)
	if err != nil {
		return nil, fmt.Errorf("module service create error: %v", err)
	}

	return modules, nil
}
