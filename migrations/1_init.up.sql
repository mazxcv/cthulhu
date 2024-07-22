CREATE TABLE IF NOT EXISTS url(
    id intger PRIMARY KEY,
    alias text NOT NULL UNIQUE,
    url text NOT NULL,
    created_at CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_alias on url(alias);