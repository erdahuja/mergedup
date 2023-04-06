-- Version: 1.01
-- Description: Create table users
CREATE TABLE users (
	user_id       UUID,
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
	item_id   UUID,
	name         TEXT,
	cost         INT,
	quantity     INT,
	user_id      UUID,
	date_created TIMESTAMP,
	date_updated TIMESTAMP,

	PRIMARY KEY (item_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Version: 1.03
-- Description: Create table cart
CREATE TABLE cart (
	cart_id      UUID,
	user_id      UUID,
	item_id   UUID,
	quantity     INT,
	paid         INT,
	date_created TIMESTAMP,

	PRIMARY KEY (cart_id),
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	FOREIGN KEY (item_id) REFERENCES items(item_id) ON DELETE CASCADE
);
