create table users(
	id uuid unique not null primary key,
	fullname varchar(100) not null,
	email varchar(150) unique not null,
	password varchar(120) not null,
	created_at timestamp default now(),
	modified_at timestamp default now()
);
