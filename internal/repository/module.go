package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"kazusa-server/internal/entity"
	"strings"
)

const (
	insertStatement = "insert into modules(id, course_id, name, content, duration_minutes) values(uuid_to_bin(?), uuid_to_bin(?), ?, ?, ?)"
	selectStatement = "select id, course_id, created_at, updated_at, name, content, duration_minutes from modules"
)

type ModuleRepositoryImplementation interface {
	Create(module entity.NewModule) (bool, error)
	Read(filters entity.ModuleFilters, pagination entity.Pagination) ([]entity.Module, error)
}

type ModuleRepository struct {
	db *sql.DB
}

func NewModuleRepo(db *sql.DB) *ModuleRepository {
	return &ModuleRepository{
		db: db,
	}
}

func (r *ModuleRepository) Create(module entity.NewModule) (bool, error) {
	newID := uuid.New()

	_, err := r.db.Exec(insertStatement, newID, module.CourseID, module.Name, module.Content, module.DurationMinutes)
	if err != nil {
		return false, fmt.Errorf("module repo error when adding new module: %v", err)
	}

	return true, nil
}

func (r *ModuleRepository) Read(filters entity.ModuleFilters, pagination entity.Pagination) ([]entity.Module, error) {
	modules := make([]entity.Module, 0, 1)

	if filters.ID == uuid.Nil && filters.CourseID == uuid.Nil {
		if pagination.Limit == 0 {
			return nil, fmt.Errorf("module repo error: at least one filter has to be passed")
		}
	}

	statement := selectStatement

	if filters.ID != uuid.Nil || filters.CourseID != uuid.Nil {
		statement += " where "
	}

	args := make([]any, 0, 4)

	if filters.ID != uuid.Nil {
		statement += "id = uuid_to_bin(?), "
		args = append(args, filters.ID)
	}

	if filters.CourseID != uuid.Nil {
		statement += "course_id = uuid_to_bin(?), "
		args = append(args, filters.CourseID)
	}

	statement = strings.TrimSuffix(statement, ", ")

	if pagination.Limit == 0 && filters.ID != uuid.Nil {
		pagination.Limit = 1
	}

	if filters.CourseID != uuid.Nil {
		pagination.Limit = 20
	}

	statement += " limit ? offset ?"

	args = append(args, pagination.Limit)
	args = append(args, pagination.Offset)

	fmt.Println("statement: ", statement)
	fmt.Println("args: ", args)

	rows, err := r.db.Query(statement, args...)
	if err != nil {
		return nil, fmt.Errorf("module repo error on reading modules: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		module := entity.Module{}

		err = rows.Scan(&module.ID, &module.CourseID, &module.CreatedAt, &module.UpdatedAt, &module.Name, &module.Content, &module.DurationMinutes)
		if err != nil {
			return nil, fmt.Errorf("module repo error on scanning a module: %v", err)
		}

		modules = append(modules, module)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("module repo error on rows when reading: %v", err)
	}

	return modules, nil
}
