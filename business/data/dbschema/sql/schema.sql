-- Version: 1.01
-- Description: Create table users
CREATE TABLE users (
	user_id       SERIAL,
	name          TEXT,
	email         TEXT UNIQUE,
	roles         TEXT[],
	password_hash TEXT,
    active        BOOLEAN,
	date_created  TIMESTAMP,
	date_updated  TIMESTAMP,

	PRIMARY KEY (user_id)
);

-- Version: 1.02
-- Description: Create table items
CREATE TABLE items (
	item_id   SERIAL,
	name         TEXT,
	cost         INT,
	quantity     INT,
	user_id      SERIAL,
	date_created TIMESTAMP,
	date_updated TIMESTAMP,

	PRIMARY KEY (item_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Version: 1.03
-- Description: Create table cart
CREATE TABLE cart (
	cart_id      SERIAL,
	user_id      SERIAL,
	item_id   SERIAL,
	quantity     INT,
	paid         INT,
	date_created TIMESTAMP,

	PRIMARY KEY (cart_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	FOREIGN KEY (item_id) REFERENCES items(item_id) ON DELETE CASCADE
);
