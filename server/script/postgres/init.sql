-- CREATE DATABASE IF NOT EXISTS a_studio SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- CREATE DATABASE a_studio;

-- enter db
\c a_studio;

-- vendor management

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  "name" VARCHAR(50),
  birthday DATE,
  account VARCHAR(50) NOT NULL UNIQUE,
  "password" VARCHAR(50) NOT NULL,
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS vendors (
  id SERIAL PRIMARY KEY,
  vendor_uuid UUID NOT NULL UNIQUE,
  db_name VARCHAR(50) NOT NULL CHECK(db_name LIKE 'a$_studio$_%' ESCAPE '$')
);

CREATE TABLE IF NOT EXISTS user_vendors (
  user_id INT REFERENCES users(id),
  vendor_id INT REFERENCES vendors(id)
);

-- customer management

CREATE TABLE IF NOT EXISTS customer (
  id SERIAL PRIMARY KEY,
  "name" VARCHAR(50),
  birthday DATE,
  account VARCHAR(50) NOT NULL UNIQUE,
  line_id VARCHAR(50) NOT NULL UNIQUE,
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS customer_members (
  customer_id INT REFERENCES customer(id),
  vendor_id INT REFERENCES vendors(id) NOT NULL
);