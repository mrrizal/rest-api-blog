CREATE TABLE IF NOT EXISTS users(
	id serial,
	username varchar(50) NOT NULL UNIQUE,
	email text NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

