package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/google/uuid"
	"strings"
	"time"
)

const (
	courseInsertStatement = "insert into courses(id, title, description, price, cover_url, attachment_urls) values(uuid_to_bin(?), ?, ?, ?, ?, ?)"
	courseSelectStatement = "select id, created_at, updated_at, title, description, price, cover_url, attachment_urls from courses"
	courseUpdateStatement = "update courses set "
	courseDeleteStatement = "delete from courses where id = uuid_to_bin(?)"
)

type CourseRepositoryImplementation interface {
	Create(course CourseCreateBody) (bool, error)
	Read(pagination entity.Pagination, filters entity.CourseFilters) ([]Course, error)
	Update(body entity.CourseUpdateBody) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

type CourseCreateBody struct {
	Title          string
	Description    string
	Price          int64
	CoverURL       string
	AttachmentURLs []string
}

func (r *CourseRepository) Create(course CourseCreateBody) (bool, error) {
	newID := uuid.New()

	attachmentsURLs := struct {
		AttachmentURLs []string `json:"attachment_urls"`
	}{
		AttachmentURLs: course.AttachmentURLs,
	}
	attachmentsURLsJSON, err := json.Marshal(attachmentsURLs)

	_, err = r.db.Exec(courseInsertStatement, newID, course.Title, course.Description, course.Price, course.CoverURL, attachmentsURLsJSON)
	if err != nil {
		return false, fmt.Errorf("course repo error when adding new course: %v", err)
	}

	return true, nil
}

type Course struct {
	ID             uuid.UUID      `db:"id"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
	Title          string         `db:"title"`
	Description    string         `db:"description"`
	Price          int64          `db:"price"`
	CoverURL       string         `db:"cover_url"`
	AttachmentURLs sql.NullString `db:"attachment_urls"`
}

func (r *CourseRepository) Read(pagination entity.Pagination, filters entity.CourseFilters) ([]Course, error) {
	courses := make([]Course, 0, pagination.Limit)

	statement := courseSelectStatement
	args := make([]any, 0, 3)

	if filters.ID != uuid.Nil {
		statement += " where "
		statement += "id = uuid_to_bin(?), "
		args = append(args, filters.ID)
	}

	if pagination.Limit == 0 {
		pagination.Limit = 1
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
		course := Course{}

		err = rows.Scan(&course.ID, &course.CreatedAt, &course.UpdatedAt, &course.Title, &course.Description, &course.Price, &course.CoverURL, &course.AttachmentURLs)
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

func (r *CourseRepository) Update(body entity.CourseUpdateBody) (bool, error) {
	statement := courseUpdateStatement
	args := make([]any, 0, 4)

	if body.ID == uuid.Nil {
		return false, fmt.Errorf("course repo error: id is empty")
	}

	if body.Title != nil {
		statement += "title = ?, "
		args = append(args, body.Title)
	}

	if body.Description != nil {
		statement += "description = ?, "
		args = append(args, body.Description)
	}

	if body.Price != nil {
		statement += "price = ?, "
		args = append(args, body.Price)
	}

	if len(args) == 0 {
		return false, fmt.Errorf("course repo error when updating course: update body is empty")
	}

	statement = strings.TrimSuffix(statement, ", ")
	args = append(args, body.ID)

	statement += " where id = uuid_to_bin(?);"

	_, err := r.db.Exec(statement, args...)
	if err != nil {
		return false, fmt.Errorf("course repo error when updating course: %v", err)
	}

	return true, nil
}

func (r *CourseRepository) Delete(id uuid.UUID) (bool, error) {
	if id == uuid.Nil {
		return false, fmt.Errorf("course repo error when deleting course: id is empty")
	}

	_, err := r.db.Exec(courseDeleteStatement, id)
	if err != nil {
		return false, fmt.Errorf("course repo error when deleting course: %v", err)
	}

	return true, nil
}
