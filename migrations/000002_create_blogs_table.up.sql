CREATE TABLE IF NOT EXISTS posts(
	id serial,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP,
	is_private BOOLEAN NOT NULL DEFAULT false,
	owner_id int NOT NULL,
	text text NOT NULL,
	title varchar(50) NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_owner
		FOREIGN KEY(owner_id)
			REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS idx_owner_id ON posts(owner_id);

CREATE INDEX IF NOT EXISTS idx_created_at ON posts(created_at DESC NULLS LAST);

