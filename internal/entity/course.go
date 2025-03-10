package entity

import (
	"github.com/google/uuid"
	"time"
)

type Course struct {
	ID             uuid.UUID `db:"id" json:"id" validate:"required"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt" validate:"required"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`
	Title          string    `db:"title" json:"title" validate:"required"`
	Description    string    `db:"description" json:"description" validate:"required"`
	Price          int64     `db:"price" json:"price" validate:"required"`
	CoverURL       string    `db:"cover_url" json:"coverUrl" validate:"required"`
	AttachmentURLs string    `db:"attachment_urls" json:"attachmentUrls"`
	Modules        *[]Module `json:"modules"`
	IsPaid         bool      `json:"isPaid"`
} // @name Course

type CourseUpdateBody struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required"`
	Title       *string   `db:"title" json:"title"`
	Description *string   `db:"description" json:"description"`
	Price       *int64    `db:"price" json:"price"`
} // @name CourseUpdateBody

type CourseFilters struct {
	ID uuid.UUID `db:"id" json:"id" validate:"required"`
} // @name CourseFilters

type CourseReadRequest struct {
	Filters    CourseFilters `json:"filters"`
	Pagination Pagination    `json:"pagination" validate:"required"`
} // @name CourseReadRequest
