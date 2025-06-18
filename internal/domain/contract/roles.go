/*

type Role struct {
	ID          int    `json:"id"`
	CinemaID    string `json:"cinema_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Permission struct {
	ID     int    `json:"id,omitempty"`
	Module string `json:"module"`
	Action string `json:"action"`
}
*/

package contract

import (
	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/google/uuid"
)

type RoleInterface interface {
	GetPermissions() ([]dto.Permission, error)
	GetPermissionsByRoleID(roleID int) ([]dto.Permission, error)
	FindByID(id int) (*dto.Role, error)
	FindByCinemaID(cinemaID uuid.UUID) ([]dto.Role, error)
	Create(role *dto.Role) (*dto.Role, error)
	Update(role *dto.Role) (*dto.Role, error)
	Delete(id int) error
	AssignPermission(roleID int, permissionID int) error
	RemovePermission(roleID int, permissionID int) error
}
