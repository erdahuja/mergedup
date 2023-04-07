INSERT INTO users (name, email, roles, password_hash, active, date_created, date_updated) VALUES
	('Merge Admin', 'admin@example.com', '{ADMIN,USER}', '$2a$10$ys4MJEHOlZr9ADxOI1xZd.WYEtxTyaE3nDgYnfDKHL/Xa77LRxVVS', true, '2023-03-24 00:00:00', '2023-03-24 00:00:00'),
	('Merge User1', 'user1@example.com', '{USER}', '$2a$10$P9QGaSITodQRPjyl8YKqQuMCA.1BFHDJhvYWmlGykfORoSJ24JYlK', true, '2023-03-24 00:00:00', '2023-03-24 00:00:00')
	ON CONFLICT DO NOTHING;

INSERT INTO items (name, cost, quantity, date_created, date_updated) VALUES
	('Comic Books', 5, 42, '2023-04-10 00:00:01.000001+00', '2023-04-10 00:00:01.000001+00'),
	( 'McDonalds Toys', 8, 120, '2023-04-10 00:00:02.000001+00', '2023-04-10 00:00:02.000001+00'),
	( 'iphone 14', 100, 2, '2023-04-10 00:00:02.000001+00', '2023-04-10 00:00:02.000001+00'),
	( 'Skoda Octavia', 550, 3, '2023-04-10 00:00:02.000001+00', '2023-04-10 00:00:02.000001+00'),
	( 'Netflix subscritption', 3, 100, '2023-04-10 00:00:02.000001+00', '2023-04-10 00:00:02.000001+00'),
	( 'Samsung AC', 75, 4, '2023-04-10 00:00:02.000001+00', '2023-04-10 00:00:02.000001+00')
	ON CONFLICT DO NOTHING;

INSERT INTO cart (user_id, date_created, date_updated) VALUES
	(1, '2023-04-10 00:00:03.000001+00', '2023-04-10 00:00:03.000001+00'),
	(2, '2023-04-10 00:00:04.000001+00', '2023-04-10 00:00:03.000001+00'),
	(2, '2023-04-10 00:00:05.000001+00', '2023-04-10 00:00:03.000001+00')
	ON CONFLICT DO NOTHING;

INSERT INTO cart_items (cart_id, item_id, quantity, date_created, date_updated) VALUES
	(1, 1, 3, '2023-04-10 00:00:03.000001+00', '2023-04-10 00:00:03.000001+00'),
	(1, 3, 4, '2023-04-10 00:00:04.000001+00', '2023-04-10 00:00:03.000001+00'),
	(2, 1, 2, '2023-04-10 00:00:05.000001+00', '2023-04-10 00:00:03.000001+00'),
	(2, 5, 1, '2023-04-10 00:00:05.000001+00', '2023-04-10 00:00:03.000001+00'),
	(3, 6, 1, '2023-04-10 00:00:05.000001+00', '2023-04-10 00:00:03.000001+00')
	ON CONFLICT DO NOTHING;