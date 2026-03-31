CREATE TABLE IF NOT EXISTS post_short_urls (
  id UUID DEFAULT uuid_generate_v7() PRIMARY KEY NOT NULL,
  full_url TEXT,
  code TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS post_short_url_id_index ON post_short_urls (id);
CREATE INDEX IF NOT EXISTS post_short_url_created_at_index ON post_short_urls (created_at);
CREATE INDEX IF NOT EXISTS post_short_url_updated_at_index ON post_short_urls (updated_at);
CREATE INDEX IF NOT EXISTS post_short_url_deleted_at_index ON post_short_urls (deleted_at);
