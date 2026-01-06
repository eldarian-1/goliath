CREATE SCHEMA IF NOT EXISTS goliath;

CREATE TABLE IF NOT EXISTS goliath.users (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    email       TEXT NOT NULL UNIQUE,
    password    TEXT NOT NULL,
    permissions TEXT[] NOT NULL DEFAULT ARRAY['read:own', 'write:own']::TEXT[],
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_users_email ON goliath.users(email) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS goliath.videos (
    id              BIGSERIAL PRIMARY KEY,
    title           TEXT NOT NULL,
    description     TEXT,
    file_name       TEXT NOT NULL,
    file_size       BIGINT NOT NULL,
    content_type    TEXT NOT NULL,
    duration        INTEGER,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_videos_created_at ON goliath.videos(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_videos_deleted_at ON goliath.videos(deleted_at) WHERE deleted_at IS NULL;
