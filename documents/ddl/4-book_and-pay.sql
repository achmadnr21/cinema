create table bookings(
	id uuid unique primary key,
	user_id uuid not null,
	schedule_id uuid not null,
	booked_at timestamp default now(),
	foreign key(user_id) references users(id),
	foreign key(schedule_id) references schedules(id)
);


CREATE TYPE seat_status AS ENUM (
  'LOCKED', 'BOOKED', 'EXPIRED'
);


create table booking_seats(
	seat_id integer not null,
	booking_id uuid not null,
	status seat_status default 'LOCKED',
	foreign key(seat_id) references seats(id),
	foreign key(booking_id) references bookings(id)
);

CREATE TYPE payment_status AS ENUM (
  'PENDING',
  'PAID',
  'FAILED',
  'REFUNDED'
);

create table payments(
	id uuid unique primary key,
	booking_id uuid unique not null,
	amount numeric(12,2) not null,
	status payment_status default 'PENDING',
	paid_at timestamp null,
	refund_reason text default 'no refund',
	refund_at timestamp null,
	created_at timestamp default now(),
	foreign key(booking_id) references bookings(id)
);