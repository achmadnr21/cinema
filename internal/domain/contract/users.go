package contract

import (
	"github.com/achmadnr21/cinema/internal/domain/dto"
)

// The Contract Interface defines the methods that the UserService must implement.

type UserInterface interface {
	// CreateUser creates a new user with the provided details.
	FindByID(id string) (*dto.User, error)
	FindByEmail(email string) (*dto.User, error)
	Create(user *dto.User) (*dto.User, error)
	Update(user *dto.User) (*dto.User, error)
	Delete(id string) error
	FindAll() ([]dto.User, error)
}
