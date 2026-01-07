-- Replace status column with progress column
ALTER TABLE goliath.videos DROP COLUMN IF EXISTS status;
ALTER TABLE goliath.videos ADD COLUMN IF NOT EXISTS progress INTEGER NOT NULL DEFAULT 0;
CREATE INDEX IF NOT EXISTS idx_videos_progress ON goliath.videos(progress) WHERE deleted_at IS NULL;

