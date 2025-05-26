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

CREATE TABLE events (
    id              BIGSERIAL PRIMARY KEY,
    participant_id  BIGINT      NOT NULL,
    user_id         BIGINT      NOT NULL,
    rrule           TEXT
);

CREATE TABLE occurrences (
    id          BIGSERIAL PRIMARY KEY,
    event_id    BIGINT       NOT NULL,
    occurs_at   TIMESTAMPTZ  NOT NULL
);