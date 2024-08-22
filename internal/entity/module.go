package entity

import (
	"github.com/google/uuid"
	"time"
)

type Module struct {
	ID              uuid.UUID `db:"id" json:"id" validate:"required"`
	CourseID        uuid.UUID `db:"course_id" json:"courseId" validate:"required"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt" validate:"required"`
	UpdatedAt       time.Time `db:"updated_at" json:"updatedAt"`
	Name            string    `db:"name" json:"name" validate:"required"`
	Content         string    `db:"content" json:"content" validate:"required"`
	DurationMinutes int64     `db:"duration_minutes" json:"durationMinutes" validate:"required"`
} // @name Module

type NewModule struct {
	CourseID        uuid.UUID `db:"course_id" json:"courseId" validate:"required"`
	Name            string    `db:"name" json:"name" validate:"required"`
	Content         string    `db:"content" json:"content" validate:"required"`
	DurationMinutes int64     `db:"duration_minutes" json:"durationMinutes" validate:"required"`
} // @name NewModule

type ModuleUpdateBody struct {
	ID              uuid.UUID `db:"id" json:"id" validate:"required"`
	Name            *string   `db:"name" json:"name"`
	Content         *string   `db:"content" json:"content"`
	DurationMinutes *int64    `db:"duration_minutes" json:"durationMinutes"`
} // @name ModuleUpdateBody

type ModuleFilters struct {
	ID       uuid.UUID `db:"id" json:"id"`
	CourseID uuid.UUID `db:"course_id" json:"courseId"`
} //@name ModuleFilters

type ModuleReadRequest struct {
	Filters    ModuleFilters `json:"filters"`
	Pagination Pagination    `json:"pagination"`
}
