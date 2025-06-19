
create table cinemas(
	id uuid unique primary key,
	name varchar(100) not null,
	address text not null
);


create table employee_details(
	cinema_id uuid not null,
	user_id uuid not null,
	role_id integer not null,
	primary key(cinema_id, user_id, role_id),
	foreign key(cinema_id) references cinemas(id),
	foreign key(user_id) references users(id),
	foreign key(role_id) references role_engine.roles(id)
);

create table halls(
	id serial primary key,
	cinema_id uuid not null,
	name varchar(120) not null,
	seat_count int default 0,
	foreign key(cinema_id) references cinemas(id)
);

create table seats(
	id serial primary key,
	hall_id integer not null,
	row char(1) not null,
	number int not null,
	foreign key(hall_id) references halls(id),
	unique(hall_id, row, number)
);

create table movies(
	id serial primary key,
	title varchar(100) not null,
	release_date date not null,
	description text default 'no desc'
);

CREATE TYPE schedule_status AS ENUM (
  'SCHEDULED', 'CANCELLED', 'POSTPONED'
);


create table schedules(
	id uuid unique primary key,
	hall_id integer not null,
	movie_id integer not null,
	show_time timestamp not null,
	price numeric(12,2) not null,
	status schedule_status default 'SCHEDULED',
	foreign key(hall_id) references halls(id),
	foreign key(movie_id) references movies(id)
);



