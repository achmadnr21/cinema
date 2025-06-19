// CREATE TYPE schedule_status AS ENUM (
//   'SCHEDULED', 'CANCELLED', 'POSTPONED'
// );

// create table schedules(
// 	id uuid unique primary key,
// 	hall_id integer not null,
// 	movie_id integer not null,
// 	show_time timestamp not null,
// 	price numeric(12,2) not null,
// 	status schedule_status default 'SCHEDULED',
// 	foreign key(hall_id) references halls(id),
// 	foreign key(movie_id) references movies(id)
// );

package dto

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID       uuid.UUID `json:"id"`
	HallID   int       `json:"hall_id"`
	MovieID  int       `json:"movie_id"`
	ShowTime time.Time `json:"show_time"`
	Price    float64   `json:"price"`
	Status   string    `json:"status"`
}
