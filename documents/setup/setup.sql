insert into cinemas(id, name, address) values
('c412854b-e64d-43c2-aadd-b23e8583dd80', 'MASTER', 'Semarang'),
('35f54cab-9012-466e-bbf9-d739712b873f', 'Gubeng CINEMA', 'Jl. Gubeng Kertajaya');

insert into role_engine.permissions(id, module, action) values
(1, 'user', 'read'),
(2, 'role', 'create'),
(3, 'role', 'read'),
(4, 'role', 'update'),
(5, 'role', 'delete'),
(6, 'cinema', 'create'),
(7, 'cinema', 'read'),
(8, 'cinema', 'update'),
(9, 'cinema', 'delete'),
(10, 'employee', 'create'),
(11, 'employee', 'read'),
(12, 'employee', 'update'),
(13, 'employee', 'delete'),
(14, 'movie', 'create'),
(15, 'movie', 'read'),
(16, 'movie', 'update'),
(17, 'movie', 'delete'),
(18, 'hall', 'create'),
(19, 'hall', 'read'),
(20, 'hall', 'update'),
(21, 'hall', 'delete'),
(22, 'seat', 'create'),
(23, 'seat', 'read'),
(24, 'seat', 'update'),
(25, 'seat', 'delete'),
(26, 'schedule', 'create'),
(27, 'schedule', 'read'),
(28, 'schedule', 'update'),
(29, 'schedule', 'delete');

insert into role_engine.roles(id, cinema_id, name) values
(1, 'c412854b-e64d-43c2-aadd-b23e8583dd80', 'SUPERADMIN');

INSERT INTO role_engine.role_permissions(role_id, permission_id) VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(1, 7),
(1, 8),
(1, 9),
(1, 10),
(1, 11),
(1, 12),
(1, 13),
(1, 14),
(1, 15),
(1, 16),
(1, 17),
(1, 18),
(1, 19),
(1, 20),
(1, 21),
(1, 22),
(1, 23),
(1, 24),
(1, 25),
(1, 26),
(1, 27),
(1, 28),
(1, 29);

select gen_random_uuid();
insert into users(id, fullname, email, password) values
('026a1382-c6ea-4401-bcb7-8244dce6710e', 'System Admin', 'achmad@gmail.com', '$2a$12$gwDVjBuHpDpNWZF/xLHEguxZaXZyiPg7RJRz.0AWiHcqLEv0cMdqK'),
('e2addb7d-5f68-4515-b9d0-dbcfa5a0c63d', 'User', 'user@gmail.com', '$2a$12$gwDVjBuHpDpNWZF/xLHEguxZaXZyiPg7RJRz.0AWiHcqLEv0cMdqK');

insert into employee_details(cinema_id, user_id, role_id) values
('35f54cab-9012-466e-bbf9-d739712b873f', '026a1382-c6ea-4401-bcb7-8244dce6710e', 1 );

insert into halls(id, cinema_id, name) values
(1, '35f54cab-9012-466e-bbf9-d739712b873f', 'Gedung I');

INSERT INTO seats(id, hall_id, row, number) VALUES
(1, 1, 'A', 1),
(2, 1, 'A', 2),
(3, 1, 'A', 3),
(4, 1, 'A', 4),
(5, 1, 'A', 5),
(6, 1, 'A', 6),
(7, 1, 'A', 7),
(8, 1, 'A', 8),
(9, 1, 'A', 9),
(10, 1, 'A', 10);

insert into movies(id, title, description, release_date) values
(1, 'Keluarga Cemara', 'Film keluarga', '2023-12-20');

