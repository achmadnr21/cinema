package usecase

import (
	"time"

	"github.com/achmadnr21/cinema/internal/domain/contract"
	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/achmadnr21/cinema/internal/utils"
	"github.com/google/uuid"
)

type ScheduleUsecase struct {
	ScheduleRepo contract.ScheduleInterface
}

func NewScheduleUsecase(scheduleRepo contract.ScheduleInterface) *ScheduleUsecase {
	return &ScheduleUsecase{
		ScheduleRepo: scheduleRepo,
	}
}

func (uc *ScheduleUsecase) GetSchedules(cinemaID uuid.UUID) ([]dto.Schedule, error) {
	schedules, err := uc.ScheduleRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (uc *ScheduleUsecase) GetScheduleByID(id uuid.UUID) (*dto.Schedule, error) {
	schedule, err := uc.ScheduleRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (uc *ScheduleUsecase) CreateSchedule(schedule *dto.Schedule) (*dto.Schedule, error) {
	schedule.ID = uuid.New()
	createdSchedule, err := uc.ScheduleRepo.Create(schedule)
	if err != nil {
		return nil, err
	}
	return createdSchedule, nil
}

func (uc *ScheduleUsecase) DeleteSchedule(id uuid.UUID) error {
	err := uc.ScheduleRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ScheduleUsecase) UpdateSchedule(schedule *dto.Schedule) (*dto.Schedule, error) {
	updatedSchedule, err := uc.ScheduleRepo.Update(schedule)
	if err != nil {
		return nil, err
	}
	return updatedSchedule, nil
}

func (uc *ScheduleUsecase) CancelSchedule(id uuid.UUID) error {
	schedule, err := uc.ScheduleRepo.FindById(id)
	if err != nil {
		return err
	}
	if schedule == nil {
		return &utils.NotFoundError{
			Message: "Schedule not found",
		}
	}

	schedule.Status = "CANCELLED"
	_, err = uc.ScheduleRepo.Update(schedule)
	if err != nil {
		return err
	}
	return nil
}

func (uc *ScheduleUsecase) PostponeSchedule(id uuid.UUID, newShowTime time.Time) (*dto.Schedule, error) {
	schedule, err := uc.ScheduleRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	if schedule == nil {
		return nil, &utils.NotFoundError{
			Message: "Schedule not found",
		}
	}

	schedule.ShowTime = newShowTime
	schedule.Status = "POSTPONED"
	updatedSchedule, err := uc.ScheduleRepo.Update(schedule)
	if err != nil {
		return nil, err
	}
	return updatedSchedule, nil
}
