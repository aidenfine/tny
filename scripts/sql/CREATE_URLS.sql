CREATE TABLE urls (
    short_url TEXT PRIMARY KEY,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    created_by TEXT NOT NULL,
    domain TEXT
);
