CREATE TABLE users (
	id bigserial NOT NULL,
	user_number text NOT NULL,
	user_name text NOT NULL,
	CONSTRAINT users_number_key UNIQUE (user_number),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE participants (
	id bigserial NOT NULL,
	participant_number text NOT NULL,
	participant_name text NOT NULL,
	user_id int4 NOT NULL,
	CONSTRAINT participants_number_key UNIQUE (user_number),
	CONSTRAINT participants_pkey PRIMARY KEY (id)
);

ALTER TABLE participants ADD CONSTRAINT participants_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;