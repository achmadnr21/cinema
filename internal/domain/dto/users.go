// create table users(
// 	id uuid unique not null primary key,
// 	fullname varchar(100) not null,
// 	email varchar(150) unique not null,
// 	password varchar(120) not null,
// 	created_at timestamp default now(),
// 	modified_at timestamp default now()
// );

package dto

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	ID          uuid.UUID `json:"id" db:"id"`
	FullName    string    `json:"fullname" db:"fullname"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"-" db:"password"`
	Created_at  time.Time `json:"created_at" db:"created_at"`
	Modified_at time.Time `json:"modified_at" db:"modified_at"`
}

// UserCreateRequest represents the request payload for creating a new user.
type UserCreateRequest struct {
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// User Update Request represents the request payload for updating an existing user.
type UserUpdateRequest struct {
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
