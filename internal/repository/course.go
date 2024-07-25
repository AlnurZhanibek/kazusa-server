package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"kazusa-server/internal/entity"
	"strings"
)

const (
	courseInsertStatement = "insert into courses(id, title, description, price) values(uuid_to_bin(?), ?, ?, ?)"
	courseSelectStatement = "select id, created_at, updated_at, title, description, price from courses"
)

type CourseRepositoryImplementation interface {
	Create(course entity.NewCourse) (bool, error)
	Read(pagination entity.Pagination, filters entity.CourseFilters) ([]entity.Course, error)
}

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) Create(course entity.NewCourse) (bool, error) {
	newID := uuid.New()

	_, err := r.db.Exec(courseInsertStatement, newID, course.Title, course.Description, course.Price)
	if err != nil {
		return false, fmt.Errorf("course repo error when adding new course: %v", err)
	}

	return true, nil
}

func (r *CourseRepository) Read(pagination entity.Pagination, filters entity.CourseFilters) ([]entity.Course, error) {
	courses := make([]entity.Course, 0, pagination.Limit)

	statement := courseSelectStatement
	args := make([]any, 0, 3)

	if filters.ID != uuid.Nil {
		statement += " where "
		statement += "id = uuid_to_bin(?), "
		args = append(args, filters.ID)
	}

	statement = strings.TrimSuffix(statement, ", ")

	statement += " limit ? offset ?"

	args = append(args, pagination.Limit)
	args = append(args, pagination.Offset)

	rows, err := r.db.Query(statement, args...)
	if err != nil {
		return nil, fmt.Errorf("course repo error on reading courses: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		course := entity.Course{}

		err = rows.Scan(&course.ID, &course.CreatedAt, &course.UpdatedAt, &course.Title, &course.Description, &course.Price)
		if err != nil {
			return nil, fmt.Errorf("course repo error on scanning a course: %v", err)
		}

		courses = append(courses, course)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("course repo error on rows when reading: %v", err)
	}

	return courses, nil
}
