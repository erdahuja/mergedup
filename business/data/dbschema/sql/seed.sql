INSERT INTO users (name, email, roles, password_hash, active, date_created, date_updated) VALUES
	('Merge Admin', 'admin@example.com', '{ADMIN,USER}', '$2a$10$ys4MJEHOlZr9ADxOI1xZd.WYEtxTyaE3nDgYnfDKHL/Xa77LRxVVS', true, '2023-03-24 00:00:00', '2023-03-24 00:00:00'),
	('Merge User1', 'user1@example.com', '{USER}', '$2a$10$P9QGaSITodQRPjyl8YKqQuMCA.1BFHDJhvYWmlGykfORoSJ24JYlK', true, '2023-03-24 00:00:00', '2023-03-24 00:00:00'),
	('Merge User2', 'user2@example.com', '{USER}', '$2a$10$AryZGMnQ.Zxv7MowxYzYqus1mPawkDQGcPhA85t5H.A.6DM4gnxL6', true, '2023-03-24 00:00:00', '2023-03-24 00:00:00')
	ON CONFLICT DO NOTHING;

INSERT INTO items (name, cost, quantity, date_created, date_updated) VALUES
	('Comic Books', 50, 42, '2023-01-01 00:00:01.000001+00', '2023-01-01 00:00:01.000001+00'),
	( 'McDonalds Toys', 75, 120, '2023-01-01 00:00:02.000001+00', '2023-01-01 00:00:02.000001+00')
	ON CONFLICT DO NOTHING;

INSERT INTO cart (cart_id, user_id, item_id, quantity, date_created) VALUES
	(1, 1, 1, 2,'2023-01-01 00:00:03.000001+00'),
	(1, 1, 1, 5, '2023-01-01 00:00:04.000001+00'),
	(2, 2, 2, 3, '2023-01-01 00:00:05.000001+00')
	ON CONFLICT DO NOTHING;
