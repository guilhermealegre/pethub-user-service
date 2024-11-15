-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied


CREATE TABLE users.users
(
    id_users         SERIAL PRIMARY KEY,
    first_name      VARCHAR(150),
    last_name       VARCHAR(150),
    avatar_photo    VARCHAR(150),
    tag             VARCHAR(150) UNIQUE,
    active          BOOLEAN DEFAULT TRUE,
    onboard_set     BOOLEAN DEFAULT FALSE,
    password_set     BOOLEAN DEFAULT FALSE,
    created_at      timestamptz NOT NULL DEFAULT now(),
    updated_at      timestamptz NOT NULL DEFAULT now()
);


-- +migrate Down
-- SQL in section 'Down' is executed when this migration is applied


