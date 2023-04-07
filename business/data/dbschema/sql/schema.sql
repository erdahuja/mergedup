-- Version: 1.01
-- Description: Create table users
CREATE TABLE users (
  id SERIAL, 
  name TEXT, 
  email TEXT UNIQUE, 
  roles TEXT[], 
  password_hash TEXT, 
  active BOOLEAN, 
  date_created TIMESTAMP, 
  date_updated TIMESTAMP, 
  PRIMARY KEY (id)
);

-- Version: 1.02
-- Description: Create table items
CREATE TABLE items (
  id SERIAL, 
  name TEXT, 
  cost INT, 
  quantity INT, 
  user_id SERIAL, 
  date_created TIMESTAMP, 
  date_updated TIMESTAMP, 
  PRIMARY KEY (id), 
  FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Version: 1.03
-- Description: Create table cart
CREATE TABLE cart (
  id SERIAL, 
  user_id INT, 
  date_created TIMESTAMP, 
  date_updated TIMESTAMP, 
  FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE, 
  );

-- Version: 1.04
-- Description: Create table cart_items
CREATE TABLE cart_items (
  id SERIAL, 
  cart_id INT,
  item_id INT, 
  quantity INT, 
  date_created TIMESTAMP, 
  date_updated TIMESTAMP, 
  PRIMARY KEY (id), 
  FOREIGN KEY (cart_id) REFERENCES cart(id) ON DELETE CASCADE, 
  FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
);
