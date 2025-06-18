package handler

import (
	"strconv"

	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/achmadnr21/cinema/internal/middleware"
	"github.com/achmadnr21/cinema/internal/usecase"
	"github.com/achmadnr21/cinema/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleHandler struct {
	uc *usecase.RoleUsecase
}

func NewRoleHandler(apiV *gin.RouterGroup, uc *usecase.RoleUsecase) {
	RoleHandler := &RoleHandler{
		uc: uc,
	}
	permission := apiV.Group("/roles")
	permission.GET("/permissions", RoleHandler.GetPermissions)

	role := apiV.Group("/cinema/:cinema_id/roles")

	role.Use(middleware.JWTAuthMiddleware)
	{
		role.POST("/",
			middleware.RequirePermissionCinema("role", "create"),
			RoleHandler.Create)
		// role.PUT("/", RoleHandler.Update)
		role.DELETE("/:id",
			middleware.RequirePermissionCinema("role", "delete"),
			RoleHandler.Delete)

		role.GET("/",
			middleware.RequirePermissionCinema("role", "read"),
			RoleHandler.FindByCinemaID)
		role.GET("/:id",
			middleware.RequirePermissionCinema("role", "read"),
			RoleHandler.FindByID)

		role.POST("/:id/permissions/:permission_id",
			middleware.RequirePermissionCinema("role", "create"),
			RoleHandler.AssignPermission)
		role.DELETE("/:id/permissions/:permission_id",
			middleware.RequirePermissionCinema("role", "create"),
			RoleHandler.RemovePermission)
		role.GET("/:id/permissions",
			middleware.RequirePermissionCinema("role", "read"), RoleHandler.GetPermissionsByRoleID)
	}
}

func (h *RoleHandler) GetPermissions(c *gin.Context) {
	permissions, err := h.uc.GetPermissions()
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Permissions retrieved successfully", permissions))
}

func (h *RoleHandler) Create(c *gin.Context) {
	cinemaIdStr := c.Param("cinema_id")
	cinema_id, err := uuid.Parse(cinemaIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid cinema ID format"))
		return
	}
	// bind the request body to the Role struct
	var req dto.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, utils.ResponseError(err.Error()))
		return
	}
	req.CinemaID = cinema_id

	role, err := h.uc.Create(&req)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(201, utils.ResponseSuccess("Role created successfully", role))
}

func (h *RoleHandler) Delete(c *gin.Context) {
	roleIdStr := c.Param("id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid role ID format"))
		return
	}

	err = h.uc.Delete(roleId)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Role deleted successfully", nil))
}

func (h *RoleHandler) FindByCinemaID(c *gin.Context) {
	cinemaIdStr := c.Param("cinema_id")
	cinema_id, err := uuid.Parse(cinemaIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid cinema ID format"))
		return
	}

	roles, err := h.uc.FindByCinemaID(cinema_id)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Roles retrieved successfully", roles))
}

func (h *RoleHandler) FindByID(c *gin.Context) {
	roleIdStr := c.Param("id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid role ID format"))
		return
	}

	role, err := h.uc.FindByID(roleId)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Role retrieved successfully", role))
}

func (h *RoleHandler) AssignPermission(c *gin.Context) {
	roleIdStr := c.Param("id")
	permissionIdStr := c.Param("permission_id")

	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid role ID format"))
		return
	}

	permissionId, err := strconv.Atoi(permissionIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid permission ID format"))
		return
	}

	err = h.uc.AssignPermission(roleId, permissionId)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Permission assigned successfully", nil))
}
func (h *RoleHandler) RemovePermission(c *gin.Context) {
	roleIdStr := c.Param("id")
	permissionIdStr := c.Param("permission_id")

	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid role ID format"))
		return
	}

	permissionId, err := strconv.Atoi(permissionIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid permission ID format"))
		return
	}

	err = h.uc.RemovePermission(roleId, permissionId)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Permission removed successfully", nil))
}

func (h *RoleHandler) GetPermissionsByRoleID(c *gin.Context) {
	roleIdStr := c.Param("id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid role ID format"))
		return
	}

	permissions, err := h.uc.GetPermissionsByRoleID(roleId)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Permissions retrieved successfully", permissions))
}
