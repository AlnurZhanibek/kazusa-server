package entity

import (
	"github.com/google/uuid"
	"time"
)

type Course struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt" validate:"required"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
	Title       string    `db:"title" json:"title" validate:"required"`
	Description string    `db:"description" json:"description" validate:"required"`
	Price       int64     `db:"price" json:"price" validate:"required"`
} // @name Course

type NewCourse struct {
	Title       string `db:"title" json:"title" validate:"required"`
	Description string `db:"description" json:"description" validate:"required"`
	Price       int64  `db:"price" json:"price" validate:"required"`
} // @name NewCourse

type CourseFilters struct {
	ID uuid.UUID `db:"id" json:"id" validate:"required"`
} // @name CourseFilters

type CourseReadRequest struct {
	Filters    CourseFilters `json:"filters"`
	Pagination Pagination    `json:"pagination" validate:"required"`
} // @name CourseReadRequest
