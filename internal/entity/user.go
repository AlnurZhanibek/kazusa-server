package entity

import (
	"github.com/google/uuid"
	"time"
)

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"createdAt" validate:"required"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Name      string    `db:"name" json:"name" validate:"required"`
	Email     string    `db:"email" json:"email" validate:"required"`
	Phone     string    `db:"phone" json:"phone" validate:"required"`
	Password  *string   `db:"password" json:"password"`
	Role      Role      `db:"role" json:"role"`
} // @name User

type NewUser struct {
	Name     string `db:"name" json:"name" validate:"required"`
	Email    string `db:"email" json:"email" validate:"required"`
	Phone    string `db:"phone" json:"phone" validate:"required"`
	Password string `db:"password" json:"password" validate:"required"`
	Role     Role   `db:"role" json:"role" validate:"required"`
} // @name NewUser

type UserUpdateBody struct {
	ID       uuid.UUID `db:"id" json:"id" validate:"required"`
	Name     *string   `db:"name" json:"name"`
	Email    *string   `db:"email" json:"email"`
	Phone    *string   `db:"phone" json:"phone"`
	Role     *Role     `db:"role" json:"role"`
	Password *string   `db:"password" json:"password"`
} // @name UserUpdateBody

type UserFilters struct {
	ID    *uuid.UUID `db:"id" json:"id"`
	Email *string    `db:"email" json:"email"`
} // @name UserFilters

type UserReadRequest struct {
	Filters    *UserFilters `json:"filters"`
	Pagination Pagination   `json:"pagination" validate:"required"`
} // @name UserReadRequest
