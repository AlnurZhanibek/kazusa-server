package repository

import (
	"database/sql"
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/google/uuid"
	"strings"
)

const (
	userInsertStatement = "insert into users(id, email, name, phone, password, role) values(uuid_to_bin(?), ?, ?, ?, ?, ?)"
	userSelectStatement = "select id, email, name, phone, password, role from users"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Phone    string
	Password string
	Role     string
}

func (r *UserRepository) Create(user User) (bool, error) {
	newID := uuid.New()

	_, err := r.db.Exec(userInsertStatement, newID, user.Email, user.Name, user.Phone, user.Password, user.Role)
	if err != nil {
		return false, fmt.Errorf("user repo error when adding new user: %v", err)
	}

	return true, nil
}

type UserFilters struct {
	ID    uuid.UUID
	Email string
}

func (r *UserRepository) Read(pagination entity.Pagination, filters UserFilters) ([]User, error) {
	users := make([]User, 0, pagination.Limit)

	statement := userSelectStatement
	args := make([]any, 0, 4)

	if filters.ID != uuid.Nil || filters.Email != "" {
		statement = statement + " where "
	}

	if filters.ID != uuid.Nil {
		statement += "id = uuid_to_bin(?), "
		args = append(args, filters.ID)
	}

	if filters.Email != "" {
		statement += "email = ?, "
		args = append(args, filters.Email)
	}

	statement = strings.TrimSuffix(statement, ", ")

	statement = statement + " limit ? offset ?"
	args = append(args, pagination.Limit)
	args = append(args, pagination.Offset)

	rows, err := r.db.Query(statement, args...)
	if err != nil {
		return nil, fmt.Errorf("user repo error on reading courses: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}

		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Phone, &user.Password, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("user repo error on scanning a user: %v", err)
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("user repo error on rows when reading: %v", err)
	}

	return users, nil
}
