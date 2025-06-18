package repository

/*
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
*/

import (
	"database/sql"
	"time"

	"github.com/achmadnr21/cinema/internal/domain/dto"
	"github.com/google/uuid"
)

type ScheduleRepository struct {
	db *sql.DB
}

func NewScheduleRepository(db *sql.DB) *ScheduleRepository {
	return &ScheduleRepository{
		db: db,
	}
}
func (r *ScheduleRepository) Create(schedule *dto.Schedule) (*dto.Schedule, error) {
	query := `
		INSERT INTO schedules (id, hall_id, movie_id, show_time, price)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.db.QueryRow(query, schedule.ID, schedule.HallID, schedule.MovieID, schedule.ShowTime, schedule.Price).Scan(&schedule.ID)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *ScheduleRepository) FindById(id uuid.UUID) (*dto.Schedule, error) {
	query := "SELECT id, hall_id, movie_id, show_time, price FROM schedules WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var schedule dto.Schedule
	err := row.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, err
	}

	return &schedule, nil
}

// ==== Read (Multiple) ===
func (r *ScheduleRepository) FindAll() ([]dto.Schedule, error) {
	query := "SELECT id, hall_id, movie_id, show_time, price FROM schedules"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []dto.Schedule
	for rows.Next() {
		var schedule dto.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}
func (r *ScheduleRepository) FindByMovieId(movieId int) ([]dto.Schedule, error) {
	query := "SELECT id, hall_id, movie_id, show_time, price FROM schedules WHERE movie_id = $1"
	rows, err := r.db.Query(query, movieId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []dto.Schedule
	for rows.Next() {
		var schedule dto.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}
func (r *ScheduleRepository) FindByCinemaID(cinemaId string) ([]dto.Schedule, error) {
	query := "SELECT id, hall_id, movie_id, show_time, price FROM schedules WHERE cinema_id = $1"
	rows, err := r.db.Query(query, cinemaId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []dto.Schedule
	for rows.Next() {
		var schedule dto.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepository) FindByShowTime(start, end time.Time) ([]dto.Schedule, error) {
	query := "SELECT id, hall_id, movie_id, show_time, price FROM schedules WHERE show_time BETWEEN $1 AND $2"
	rows, err := r.db.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []dto.Schedule
	for rows.Next() {
		var schedule dto.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

// ==== Update ===
func (r *ScheduleRepository) Update(schedule *dto.Schedule) (*dto.Schedule, error) {
	query := `
		UPDATE schedules
		SET hall_id = $1, movie_id = $2, show_time = $3, price = $4
		WHERE id = $5
	`
	_, err := r.db.Exec(query, schedule.HallID, schedule.MovieID, schedule.ShowTime, schedule.Price, schedule.ID)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}
func (r *ScheduleRepository) Delete(id uuid.UUID) error {
	query := "DELETE FROM schedules WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
