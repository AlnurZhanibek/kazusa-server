package service

import (
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/google/uuid"
	"os"
)

type ActivityService struct {
	repo         *repository.ActivityRepository
	emailService *EmailService
}

func NewActivityService(repo *repository.ActivityRepository, emailService *EmailService) *ActivityService {
	return &ActivityService{repo: repo, emailService: emailService}
}

type ActivityCreateBody struct {
	UserID       uuid.UUID
	UserEmail    string
	UserFullname string
	CourseID     uuid.UUID
	CourseName   string
	ModuleID     uuid.UUID
	IsLast       *bool
}

func (s *ActivityService) Create(activity *ActivityCreateBody) error {
	if activity.IsLast != nil && *activity.IsLast {
		_ = s.emailService.SendEmail(os.Getenv("OWNER_EMAIL"), os.Getenv("OWNER_EMAIL"),
			"Course completion", fmt.Sprintf("%v (%v) completed %v!",
				activity.UserEmail, activity.UserFullname, activity.CourseName))
	}

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
