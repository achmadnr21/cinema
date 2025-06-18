package repository

import (
	"database/sql"

	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/google/uuid"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

/*
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
*/

func (r *RoleRepository) GetPermissions() ([]dto.Permission, error) {
	query := "SELECT id, module, action FROM role_engine.permissions"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []dto.Permission
	for rows.Next() {
		var permission dto.Permission
		if err := rows.Scan(&permission.ID, &permission.Module, &permission.Action); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *RoleRepository) GetPermissionsByRoleID(roleID int) ([]dto.Permission, error) {
	query := `
		SELECT p.id, p.module, p.action
		FROM role_engine.permissions p
		JOIN role_engine.role_permissions rp ON p.id = rp.permission_id
		WHERE rp.role_id = $1`
	rows, err := r.db.Query(query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var permissions []dto.Permission
	for rows.Next() {
		var permission dto.Permission
		if err := rows.Scan(&permission.ID, &permission.Module, &permission.Action); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *RoleRepository) FindByID(id int) (*dto.Role, error) {
	query := "SELECT id, cinema_id, name, description FROM role_engine.roles WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var role dto.Role
	err := row.Scan(&role.ID, &role.CinemaID, &role.Name, &role.Description)
	if err != nil {
		return nil, err // Other error
	}

	return &role, nil
}
func (r *RoleRepository) FindByCinemaID(cinemaID uuid.UUID) ([]dto.Role, error) {
	query := "SELECT id, cinema_id, name, description FROM role_engine.roles WHERE cinema_id = $1"
	rows, err := r.db.Query(query, cinemaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []dto.Role
	for rows.Next() {
		var role dto.Role
		if err := rows.Scan(&role.ID, &role.CinemaID, &role.Name, &role.Description); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
func (r *RoleRepository) Create(role *dto.Role) (*dto.Role, error) {
	query := "INSERT INTO role_engine.roles (cinema_id, name, description) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, role.CinemaID, role.Name, role.Description).Scan(&role.ID)
	if err != nil {
		return nil, err
	}
	return role, nil
}
func (r *RoleRepository) Update(role *dto.Role) (*dto.Role, error) {
	query := "UPDATE role_engine.roles SET cinema_id = $1, name = $2, description = $3 WHERE id = $4"
	_, err := r.db.Exec(query, role.CinemaID, role.Name, role.Description, role.ID)
	if err != nil {
		return nil, err
	}
	return role, nil
}
func (r *RoleRepository) Delete(id int) error {
	query := "DELETE FROM role_engine.roles WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *RoleRepository) AssignPermission(roleID int, permissionID int) error {
	query := "INSERT INTO role_engine.role_permissions (role_id, permission_id) VALUES ($1, $2)"
	_, err := r.db.Exec(query, roleID, permissionID)
	if err != nil {
		return err
	}
	return nil
}
func (r *RoleRepository) RemovePermission(roleID int, permissionID int) error {
	query := "DELETE FROM role_engine.role_permissions WHERE role_id = $1 AND permission_id = $2"
	_, err := r.db.Exec(query, roleID, permissionID)
	if err != nil {
		return err
	}
	return nil
}
