package entity

import (
	"github.com/google/uuid"
	"time"
)

type Course struct {
	ID          uuid.UUID `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Price       int64     `db:"price"`
}

type NewCourse struct {
	Title       string `db:"title"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
}
