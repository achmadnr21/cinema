package contract

import (
	"time"

	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/google/uuid"
)

//	type Schedule struct {
//		ID       uuid.UUID `json:"id"`
//		HallID   int       `json:"hall_id"`
//		MovieID  int       `json:"movie_id"`
//		ShowTime time.Time `json:"show_time"`
//		Price    float64   `json:"price"`
//	}

type ScheduleInterface interface {
	// === Create ===
	Create(schedule *dto.Schedule) (*dto.Schedule, error)

	// === Read (Single) ===
	FindById(id uuid.UUID) (*dto.Schedule, error)

	// === Read (Multiple) ===
	FindAll() ([]dto.Schedule, error)
	FindByCinemaID(cinemaId string) ([]dto.Schedule, error)
	FindByMovieId(movieId int) ([]dto.Schedule, error)
	FindByShowTime(start, end time.Time) ([]dto.Schedule, error)

	// === Update ===
	Update(schedule *dto.Schedule) (*dto.Schedule, error)

	// === Delete ===
	Delete(id uuid.UUID) error
}
