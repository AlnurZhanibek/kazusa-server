package entity

import (
	"github.com/google/uuid"
	"time"
)

type Module struct {
	ID              uuid.UUID `db:"id" json:"id"`
	CourseID        uuid.UUID `db:"course_id" json:"courseId"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time `db:"updated_at" json:"updatedAt"`
	Name            string    `db:"name" json:"name"`
	Content         string    `db:"content" json:"content"`
	DurationMinutes int64     `db:"duration_minutes" json:"durationMinutes"`
}

type NewModule struct {
	CourseID        uuid.UUID `db:"course_id" json:"courseId"`
	Name            string    `db:"name" json:"name"`
	Content         string    `db:"content" json:"content"`
	DurationMinutes int64     `db:"duration_minutes" json:"durationMinutes"`
}

type ModuleFilters struct {
	ID       uuid.UUID `db:"id" json:"id"`
	CourseID uuid.UUID `db:"course_id" json:"courseId"`
}

type ModuleReadRequest struct {
	Filters    ModuleFilters `json:"filters"`
	Pagination Pagination    `json:"pagination"`
}
