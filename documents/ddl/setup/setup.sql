insert into cinemas(id, name, address) values
('c412854b-e64d-43c2-aadd-b23e8583dd80', 'MASTER', 'Semarang'),
('35f54cab-9012-466e-bbf9-d739712b873f', 'Gubeng CINEMA', 'Jl. Gubeng Kertajaya');

insert into role_engine.permissions(id, module, action) values
(1, 'admin', 'admin'),
(2, 'user', 'read'),
(3, 'role', 'create'),
(4, 'role', 'read'),
(5, 'role', 'update'),
(6, 'role', 'delete'),
(7, 'cinema', 'create'),
(8, 'cinema', 'read'),
(9, 'cinema', 'update'),
(10, 'cinema', 'delete'),
(11, 'employee', 'create'),
(12, 'employee', 'read'),
(13, 'employee', 'update'),
(14, 'employee', 'delete'),
(15, 'movie', 'create'),
(16, 'movie', 'read'),
(17, 'movie', 'update'),
(18, 'movie', 'delete'),
(19, 'hall', 'create'),
(20, 'hall', 'read'),
(21, 'hall', 'update'),
(22, 'hall', 'delete'),
(23, 'seat', 'create'),
(24, 'seat', 'read'),
(25, 'seat', 'update'),
(26, 'seat', 'delete'),
(27, 'schedule', 'create'),
(28, 'schedule', 'read'),
(29, 'schedule', 'update'),
(30, 'schedule', 'delete');

delete from role_engine.permissions;

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
(1, 29),
(1, 30);

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

--select * from cinemas;
--select * from users;
--select ed.cinema_id, p.module, p.action
--from employee_details ed
--inner join role_engine.roles rer on ed.role_id = rer.id
--left join role_engine.role_permissions rp on rer.id = rp.role_id 
--left join role_engine.permissions p on rp.permission_id = p.id
--where ed.cinema_id = 'c412854b-e64d-43c2-aadd-b23e8583dd80'
--;
--select * from role_engine.roles r ;
--
--select * from  employee_details ed ;
--
--
---- membuat schedule
--select * from halls;
--select * from movies;
--select * from schedules;
--
--SELECT sc.id, sc.hall_id, sc.movie_id, sc.show_time, sc.price, sc.status
--FROM schedules sc
--inner join halls h on sc.hall_id = h.id 
--WHERE h.cinema_id = '35f54cab-9012-466e-bbf9-d739712b873f';
--
--
--delete from schedules ;
