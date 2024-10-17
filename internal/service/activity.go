package service

import (
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/google/uuid"
)

type ActivityService struct {
	repo *repository.ActivityRepository
}

func NewActivityService(repo *repository.ActivityRepository) *ActivityService {
	return &ActivityService{repo: repo}
}

type ActivityCreateBody struct {
	UserID   uuid.UUID
	CourseID uuid.UUID
	ModuleID uuid.UUID
}

func (s *ActivityService) Create(activity *ActivityCreateBody) error {
	return s.repo.Create(&repository.ActivityCreateBody{
		UserID:   activity.UserID,
		CourseID: activity.CourseID,
		ModuleID: activity.ModuleID,
	})
}

type Activity struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	CourseID uuid.UUID
	ModuleID uuid.UUID
} // @name Activity

type ActivityFilters struct {
	UserID   *uuid.UUID
	CourseID *uuid.UUID
	ModuleID *uuid.UUID
}

func (s *ActivityService) Read(filters ActivityFilters) ([]Activity, error) {
	repoActivities, err := s.repo.Read(&repository.ActivityFilters{
		UserID:   filters.UserID,
		CourseID: filters.CourseID,
		ModuleID: filters.ModuleID,
	})

	if err != nil {
		return nil, fmt.Errorf("service error when reading activity: %v", err)
	}

	activities := make([]Activity, 0, len(repoActivities))
	for _, act := range repoActivities {
		activities = append(activities, Activity{
			ID:       act.ID,
			UserID:   act.UserID,
			CourseID: act.CourseID,
			ModuleID: act.ModuleID,
		})
	}

	return activities, nil
}

type Payment struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	CourseID uuid.UUID
} // @name Payment

type PaymentFilters struct {
	UserID   *uuid.UUID
	CourseID *uuid.UUID
}

func (s *PaymentService) Read(filters PaymentFilters) (*Payment, error) {
	repoPayment, err := s.repo.Read(&repository.PaymentFilters{
		UserID:   filters.UserID,
		CourseID: filters.CourseID,
	})

	if err != nil {
		return nil, fmt.Errorf("service error when reading activity: %v", err)
	}

	payment := &Payment{
		ID:       repoPayment.ID,
		UserID:   repoPayment.UserID,
		CourseID: repoPayment.CourseID,
	}

	return payment, nil
}
