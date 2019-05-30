CREATE DATABASE lenslocked_dev;

\c lenslocked_dev

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT,
  email TEXT NOT NULL
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  amount INT,
  description TEXT
);

SELECT * FROM users;

SELECT * FROM orders;

INSERT INTO users (age, email, first_name, last_name)
VALUES (30, 'jon@calhoun.io', 'Jonathan', 'Calhoun')
