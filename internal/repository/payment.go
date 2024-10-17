package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

const (
	PAYMENT_INSERT_STATEMENT = "insert into course_payments(id, user_id, course_id) values(uuid_to_bin(?), uuid_to_bin(?), uuid_to_bin(?))"
	PAYMENT_SELECT_STATEMENT = "select id, user_id, course_id from course_payments"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

type PaymentCreateBody struct {
	UserID   uuid.UUID `db:"user_id"`
	CourseID uuid.UUID `db:"course_id"`
}

func (r *PaymentRepository) Create(payment *PaymentCreateBody) error {
	newID := uuid.New()

	_, err := r.db.Exec(PAYMENT_INSERT_STATEMENT, newID, payment.UserID, payment.CourseID)
	if err != nil {
		return fmt.Errorf("payment repo error when adding new course: %v", err)
	}

	return nil
}

type Payment struct {
	ID       uuid.UUID `db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	CourseID uuid.UUID `db:"course_id"`
}

type PaymentFilters struct {
	UserID   *uuid.UUID
	CourseID *uuid.UUID
}

func (r *PaymentRepository) Read(filters *PaymentFilters) (*Payment, error) {
	if filters == nil || (filters.UserID == nil || filters.CourseID == nil) {
		return nil, fmt.Errorf("payment repo error on read: user_id and course_id filter should be passed")
	}

	payment := Payment{}

	statement := PAYMENT_SELECT_STATEMENT
	statement += " where "
	args := make([]any, 0, 2)

	if filters.UserID != nil {
		statement += "user_id = uuid_to_bin(?) and "
		args = append(args, *filters.UserID)
	}
	if filters.CourseID != nil {
		statement += "course_id = uuid_to_bin(?) and "
		args = append(args, *filters.CourseID)
	}
	statement = strings.TrimSuffix(statement, " and ")

	row := r.db.QueryRow(statement, args...)
	err := row.Scan(&payment.ID, &payment.UserID, &payment.CourseID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("payment repo error on read: %v", err)
	}

	return &payment, nil
}
