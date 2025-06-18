package handler

import (
	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/achmadnr21/cinema/internal/middleware"
	"github.com/achmadnr21/cinema/internal/usecase"
	"github.com/achmadnr21/cinema/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ScheduleHandler struct {
	uc *usecase.ScheduleUsecase
}

func NewScheduleHandler(apiV *gin.RouterGroup, uc *usecase.ScheduleUsecase) {
	scheduleHandler := &ScheduleHandler{
		uc: uc,
	}

	schedule := apiV.Group("/cinema/:cinema_id/schedules")
	schedule.Use(middleware.JWTAuthMiddleware)
	{
		schedule.GET("/", // all
			middleware.RequirePermissionCinema("schedule", "read"),
			scheduleHandler.GetSchedules)

		schedule.POST("/",
			middleware.RequirePermissionCinema("schedule", "create"),
			scheduleHandler.CreateSchedule)

		schedule.PUT("/:id",
			middleware.RequirePermissionCinema("schedule", "update"),
			scheduleHandler.UpdateSchedule)

		schedule.DELETE("/:id",
			middleware.RequirePermissionCinema("schedule", "delete"),
			scheduleHandler.DeleteSchedule)
	}
}

func (h *ScheduleHandler) GetSchedules(c *gin.Context) {

	cinemaID, err := uuid.Parse(c.Param("cinema_id"))
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid cinema ID"))
		return
	}
	schedules, err := h.uc.GetSchedules(cinemaID)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Data Retrieved", schedules))
}

func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var schedule *dto.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(400, utils.ResponseError("Invalid request body"))
		return
	}
	schedule.ID = uuid.New()
	// create it now
	schedule, err := h.uc.CreateSchedule(schedule)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}

	c.JSON(201, utils.ResponseSuccess("Schedule created successfully", schedule))
}

func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	scheduleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ResponseError("Invalid schedule ID"))
		return
	}

	schedule := new(dto.Schedule)
	if err := c.ShouldBindJSON(schedule); err != nil {
		c.JSON(400, utils.ResponseError("Invalid request body"))
		return
	}

	schedule.ID = scheduleID

	schedule, err = h.uc.UpdateSchedule(schedule)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}

	c.JSON(200, utils.ResponseSuccess("Schedule updated successfully", schedule))
}

func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	scheduleID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	err = h.uc.DeleteSchedule(scheduleID)
	if err != nil {
		c.JSON(utils.GetHTTPErrorCode(err), utils.ResponseError(err.Error()))
		return
	}
	c.JSON(200, utils.ResponseSuccess("Schedule deleted successfully", nil))
}
