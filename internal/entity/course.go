package entity

import (
	"github.com/google/uuid"
	"time"
)

type Course struct {
	ID          uuid.UUID `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"updatedAt"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Price       int64     `db:"price" json:"price"`
}

type NewCourse struct {
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Price       int64  `db:"price" json:"price"`
}

type CourseFilters struct {
	ID uuid.UUID `db:"id" json:"id"`
}

type CourseReadRequest struct {
	Filters    CourseFilters `json:"filters"`
	Pagination Pagination    `json:"pagination"`
}
