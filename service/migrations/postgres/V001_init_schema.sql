CREATE SCHEMA IF NOT EXISTS goliath;

CREATE TABLE IF NOT EXISTS goliath.users (
    id          BIGSERIAL,
    name        TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);
