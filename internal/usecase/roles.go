package usecase

import (
	"github.com/achmadnr21/cinema/internal/domain/contract"
	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/achmadnr21/cinema/internal/utils"
	"github.com/google/uuid"
)

type RoleUsecase struct {
	RoleRepo contract.RoleInterface
}

func NewRoleUsecase(RoleRepo contract.RoleInterface) *RoleUsecase {
	return &RoleUsecase{
		RoleRepo: RoleRepo,
	}
}
func (uc *RoleUsecase) GetPermissions() ([]dto.Permission, error) {
	permissions, err := uc.RoleRepo.GetPermissions()
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (uc *RoleUsecase) Create(role *dto.Role) (*dto.Role, error) {
	// langsung buat role baru
	createdRole, err := uc.RoleRepo.Create(role)
	if err != nil {
		return nil, err
	}
	return createdRole, nil
}

func (uc *RoleUsecase) Delete(roleID int) error {
	err := uc.RoleRepo.Delete(roleID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *RoleUsecase) FindByCinemaID(cinemaID uuid.UUID) ([]dto.Role, error) {
	roles, err := uc.RoleRepo.FindByCinemaID(cinemaID)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (uc *RoleUsecase) FindByID(roleID int) (*dto.Role, error) {
	role, err := uc.RoleRepo.FindByID(roleID)
	if err != nil {
		return nil, &utils.NotFoundError{
			Message: "Role not found"}
	}
	return role, nil
}
func (uc *RoleUsecase) AssignPermission(roleID int, permissionID int) error {
	err := uc.RoleRepo.AssignPermission(roleID, permissionID)
	if err != nil {
		return err
	}
	return nil
}
func (uc *RoleUsecase) RemovePermission(roleID int, permissionID int) error {
	err := uc.RoleRepo.RemovePermission(roleID, permissionID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *RoleUsecase) GetPermissionsByRoleID(roleID int) ([]dto.Permission, error) {
	permissions, err := uc.RoleRepo.GetPermissionsByRoleID(roleID)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
