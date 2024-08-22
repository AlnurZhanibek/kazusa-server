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
	userSelectStatement = "select id, email, name, phone, password, role from users;"
	userUpdateStatement = "update users set "
	userDeleteStatement = "delete from users where id = uuid_to_bin(?);"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user entity.NewUser) (bool, error) {
	newID := uuid.New()

	_, err := r.db.Exec(userInsertStatement, newID, user.Email, user.Name, user.Phone, user.Password, user.Role)
	if err != nil {
		return false, fmt.Errorf("user repo error when adding new user: %v", err)
	}

	return true, nil
}

func (r *UserRepository) Read(pagination entity.Pagination, filters entity.UserFilters) ([]entity.User, error) {
	users := make([]entity.User, 0, pagination.Limit)

	statement := userSelectStatement
	args := make([]any, 0, 4)

	if filters.ID != nil || filters.Email != nil {
		statement = statement + " where "
	}

	if filters.ID != nil {
		statement += "id = uuid_to_bin(?), "
		args = append(args, *filters.ID)
	}

	if filters.Email != nil {
		statement += "email = ?, "
		args = append(args, *filters.Email)
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
		user := entity.User{}

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

func (r *UserRepository) Update(body entity.UserUpdateBody) (bool, error) {
	statement := userUpdateStatement
	args := make([]any, 0, 6)

	if body.ID == uuid.Nil {
		return false, fmt.Errorf("ID is empty, repo layer error")
	}

	if body.Name != nil {
		statement += "name = ?, "
		args = append(args, body.Name)
	}

	if body.Email != nil {
		statement += "email = ?, "
		args = append(args, body.Email)
	}

	if body.Phone != nil {
		statement += "phone = ?, "
		args = append(args, body.Phone)
	}

	if body.Password != nil {
		statement += "password = ?, "
		args = append(args, body.Password)
	}

	if body.Role != nil {
		statement += "role = ?, "
		args = append(args, body.Role)
	}

	if len(args) == 0 {
		return false, fmt.Errorf("user repo error: update body is empty")
	}

	statement = strings.TrimSuffix(statement, ", ")
	args = append(args, body.ID)

	statement += " where id = uuid_to_bin(?);"

	_, err := r.db.Exec(statement, args...)
	if err != nil {
		return false, fmt.Errorf("user repo error when updating user: %v", err)
	}

	return true, nil
}

func (r *UserRepository) Delete(id uuid.UUID) (bool, error) {
	if id == uuid.Nil {
		return false, fmt.Errorf("user repo error when deleting user: id is empty")
	}

	_, err := r.db.Exec(userDeleteStatement, id)
	if err != nil {
		return false, fmt.Errorf("user repo error when deleting user: %v", err)
	}

	return true, nil
}
