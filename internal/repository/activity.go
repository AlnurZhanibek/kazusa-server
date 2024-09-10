package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

const (
	INSERT_STATEMENT = "insert into user_activity(id, user_id, course_id, module_id) values(uuid_to_bin(?), uuid_to_bin(?), uuid_to_bin(?), uuid_to_bin(?))"
	SELECT_STATEMENT = "select id, user_id, course_id, module_id from user_activity"
)

type ActivityRepository struct {
	db *sql.DB
}

func NewActivityRepository(db *sql.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

type ActivityCreateBody struct {
	UserID   uuid.UUID `db:"user_id"`
	CourseID uuid.UUID `db:"course_id"`
	ModuleID uuid.UUID `db:"module_id"`
}

func (r *ActivityRepository) Create(activity *ActivityCreateBody) error {
	newID := uuid.New()

	_, err := r.db.Exec(INSERT_STATEMENT, newID, activity.UserID, activity.CourseID, activity.ModuleID)
	if err != nil {
		return fmt.Errorf("activity repo error when adding new course: %v", err)
	}

	return nil
}

type Activity struct {
	ID       uuid.UUID `db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	CourseID uuid.UUID `db:"course_id"`
	ModuleID uuid.UUID `db:"module_id"`
}

type ActivityFilters struct {
	UserID   *uuid.UUID
	CourseID *uuid.UUID
	ModuleID *uuid.UUID
}

func (r *ActivityRepository) Read(filters *ActivityFilters) ([]*Activity, error) {
	if filters == nil && (filters.UserID == nil && filters.CourseID == nil && filters.ModuleID == nil) {
		return nil, fmt.Errorf("activity repo error on read: at least one filter should be passed")
	}

	activities := make([]*Activity, 0)

	statement := SELECT_STATEMENT
	statement += " where "
	args := make([]any, 0, 3)

	if filters.UserID != nil {
		statement += "user_id = uuid_to_bin(?) and "
		args = append(args, *filters.UserID)
	}
	if filters.CourseID != nil {
		statement += "course_id = uuid_to_bin(?) and "
		args = append(args, *filters.CourseID)
	}
	if filters.ModuleID != nil {
		statement += "module_id = uuid_to_bin(?) and "
		args = append(args, *filters.ModuleID)
	}
	statement = strings.TrimSuffix(statement, " and ")

	rows, err := r.db.Query(statement, args...)
	if err != nil {
		return nil, fmt.Errorf("activity repo error on read: query method")
	}
	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		activity := &Activity{}
		err = rows.Scan(&activity.ID, &activity.UserID, &activity.CourseID, &activity.ModuleID)
		if err != nil {
			return nil, fmt.Errorf("activity repo error on read: scan method")
		}
		activities = append(activities, activity)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("activity repo error on read: err return by rows")
	}

	return activities, nil
}
