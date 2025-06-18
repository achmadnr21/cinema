/*
type UserInterface interface {
	// CreateUser creates a new user with the provided details.
	FindByID(id string) (*dto.User, error)
	FindByEmail(email string) (*dto.User, error)
	Create(user *dto.User) (*dto.User, error)
	Update(user *dto.User) (*dto.User, error)
	Delete(id string) error
	FindAll() ([]dto.User, error)
}
*/

package repository

import (
	"database/sql"
	"fmt"

	"github.com/achmadnr21/cinema/internal/domain/dto"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByID(id string) (*dto.User, error) {
	query := "SELECT id, fullname, email, password FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user dto.User
	err := row.Scan(&user.ID, &user.FullName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*dto.User, error) {
	query := "SELECT id, fullname, email, password FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	var user dto.User
	err := row.Scan(&user.ID, &user.FullName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}

	return &user, nil
}
func (r *UserRepository) Create(user *dto.User) (*dto.User, error) {
	query := "INSERT INTO users (id, fullname, email, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, user.ID, user.FullName, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	return user, nil
}
func (r *UserRepository) Update(user *dto.User) (*dto.User, error) {
	query := "UPDATE users SET fullname = $1, email = $2, password = $3, modified_at = now() WHERE id = $4"
	_, err := r.db.Exec(query, user.FullName, user.Email, user.Password, user.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return user, nil
}
func (r *UserRepository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
func (r *UserRepository) FindAll() ([]dto.User, error) {
	query := "SELECT id, fullname, email, password FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving users: %v", err)
	}
	defer rows.Close()

	var users []dto.User
	for rows.Next() {
		var user dto.User
		err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %v", err)
	}

	return users, nil
}
