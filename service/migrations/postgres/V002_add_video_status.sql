ALTER TABLE goliath.videos ADD COLUMN IF NOT EXISTS status TEXT NOT NULL DEFAULT 'processing';
CREATE INDEX IF NOT EXISTS idx_videos_status ON goliath.videos(status) WHERE deleted_at IS NULL;

