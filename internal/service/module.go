package service

import (
	"context"
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/google/uuid"
)

type ContextKey string

const userIdCtxKey ContextKey = "userId"

type ModuleServiceImplementation interface {
	Create(Module entity.NewModule) (bool, error)
	Read(ctx context.Context, pagination entity.Pagination, filters entity.ModuleFilters) ([]entity.Module, error)
	Update(body entity.ModuleUpdateBody) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}

type ModuleService struct {
	repo            repository.ModuleRepositoryImplementation
	activityService *ActivityService
}

func NewModuleService(repo repository.ModuleRepositoryImplementation, activityService *ActivityService) *ModuleService {
	return &ModuleService{repo: repo, activityService: activityService}
}

func (s *ModuleService) Create(Module entity.NewModule) (bool, error) {
	ok, err := s.repo.Create(Module)
	if err != nil {
		return false, fmt.Errorf("module service create error: %v", err)
	}

	return ok, nil
}

func (s *ModuleService) Read(ctx context.Context, pagination entity.Pagination, filters entity.ModuleFilters) ([]entity.Module, error) {
	modules, err := s.repo.Read(filters, pagination)
	if err != nil {
		return nil, fmt.Errorf("module service read error: %v", err)
	}

	userIDCtx := ctx.Value("user_id")

	if filters.CourseID != uuid.Nil && userIDCtx != nil {

		userID := userIDCtx.(uuid.UUID)

		activities, actErr := s.activityService.Read(ActivityFilters{
			UserID:   &userID,
			CourseID: &filters.CourseID,
		})

		if actErr != nil {
			return nil, fmt.Errorf("module service read error when adding activity: %v", actErr)
		}

		for _, module := range modules {
			for _, activity := range activities {
				if module.ID == activity.ModuleID {
					module.IsCompleted = true
				}
			}
		}
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
