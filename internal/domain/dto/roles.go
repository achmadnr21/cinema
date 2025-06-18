package dto

import "github.com/google/uuid"

/*
set time zone 'ASIA/Jakarta';

create schema role_engine;

create table role_engine.permissions(

	id Serial primary key,
	module varchar(20) not null,
	action varchar(50) not null

);

create table role_engine.roles(

	id Serial primary key,
	cinema_id uuid not null,
	name varchar(50) not null,
	description text default 'no description'

);

create table role_engine.role_permissions(

	role_id int not null,
	permission_id int not null,
	primary key(role_id, permission_id),
	foreign key(role_id) references role_engine.roles(id) on delete cascade,
	foreign key(permission_id) references role_engine.permissions(id) on delete cascade

);
*/
type Role struct {
	ID          int       `json:"id"`
	CinemaID    uuid.UUID `json:"cinema_id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
}
type Permission struct {
	ID     int    `json:"id,omitempty"`
	Module string `json:"module"`
	Action string `json:"action"`
}
