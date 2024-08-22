package service

import (
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/google/uuid"
)

type UserServiceImplementation interface {
	Create(user entity.NewUser) (bool, error)
	Read(pagination entity.Pagination, filters entity.UserFilters) ([]entity.User, error)
	Update(body entity.UserUpdateBody) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user entity.NewUser) (bool, error) {
	ok, err := s.repo.Create(user)
	if err != nil {
		return false, fmt.Errorf("user service create error: %v", err)
	}

	return ok, nil
}

func (s *UserService) Read(pagination entity.Pagination, filters entity.UserFilters) ([]entity.User, error) {
	modules, err := s.repo.Read(pagination, filters)
	if err != nil {
		return nil, fmt.Errorf("user service create error: %v", err)
	}

	return modules, nil
}

func (s *UserService) Update(body entity.UserUpdateBody) (bool, error) {
	ok, err := s.repo.Update(body)
	if err != nil {
		return false, fmt.Errorf("user service update error: %v", err)
	}

	return ok, nil
}

func (s *UserService) Delete(id uuid.UUID) (bool, error) {
	ok, err := s.repo.Delete(id)
	if err != nil {
		return false, fmt.Errorf("user service delete error: %v", err)
	}

	return ok, nil
}
