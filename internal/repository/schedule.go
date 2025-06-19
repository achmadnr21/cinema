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
	"fmt"
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
	query := "SELECT id, hall_id, movie_id, show_time, price, status FROM schedules WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var schedule dto.Schedule
	err := row.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price, &schedule.Status)
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
	query := "SELECT id, hall_id, movie_id, show_time, price, status FROM schedules"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []dto.Schedule
	for rows.Next() {
		var schedule dto.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price, &schedule.Status); err != nil {
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
	query := "SELECT id, hall_id, movie_id, show_time, price, status FROM schedules WHERE movie_id = $1"
	rows, err := r.db.Query(query, movieId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []dto.Schedule
	for rows.Next() {
		var schedule dto.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price, &schedule.Status); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}
func (r *ScheduleRepository) FindByCinemaID(cinemaId uuid.UUID) ([]dto.Schedule, error) {
	query := `SELECT sc.id, sc.hall_id, sc.movie_id, sc.show_time, sc.price, sc.status
		FROM schedules sc
		inner join halls h on sc.hall_id = h.id 
		WHERE h.cinema_id = $1`
	rows, err := r.db.Query(query, cinemaId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []dto.Schedule
	for rows.Next() {
		var schedule dto.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price, &schedule.Status); err != nil {
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
	query := "SELECT id, hall_id, movie_id, show_time, price, status FROM schedules WHERE show_time BETWEEN $1 AND $2"
	rows, err := r.db.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []dto.Schedule
	for rows.Next() {
		var schedule dto.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.HallID, &schedule.MovieID, &schedule.ShowTime, &schedule.Price, &schedule.Status); err != nil {
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
	// buat dinamis
	sets := []string{}
	if !schedule.ShowTime.IsZero() {
		timeStr := schedule.ShowTime.Format("2006-01-02 15:04:05")
		sets = append(sets, fmt.Sprintf("show_time = '%s'", timeStr))
	}
	if schedule.Price != 0 {
		sets = append(sets, fmt.Sprintf("price = %f", schedule.Price))
	}
	if len(sets) == 0 {
		return nil, sql.ErrNoRows // No fields to update
	}
	// Join the sets with commas
	setsStr := ""
	for i, set := range sets {
		if i > 0 {
			setsStr += ", "
		}
		setsStr += set
	}
	// Prepare the query
	// Note: The status field is not included in the update query, assuming it is not being updated.
	query := "UPDATE schedules SET " + setsStr + " WHERE id = $1"

	_, err := r.db.Exec(query, schedule.ID)
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
