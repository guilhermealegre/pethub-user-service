-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied


CREATE TABLE "user".users
(
    id_users         SERIAL PRIMARY KEY,
    uuid            UUID NOT NULL UNIQUE,
    first_name      VARCHAR(150),
    last_name       VARCHAR(150),
    avatar_photo    VARCHAR(150),
    tag             VARCHAR(150) UNIQUE,
    active          BOOLEAN,
    onboard_set     BOOLEAN DEFAULT FALSE,
    created_at      timestamptz NOT NULL DEFAULT now(),
    updated_at      timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX idx_users_uuid ON "user".users(uuid);


-- +migrate Down
-- SQL in section 'Down' is executed when this migration is applied

DROP TABLE  "user".users;

