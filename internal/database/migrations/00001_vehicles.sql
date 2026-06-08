-- +goose Up
CREATE TABLE IF NOT EXISTS vehicles (
    id SERIAL PRIMARY KEY,
    brand TEXT NOT NULL DEFAULT '',
    model TEXT NOT NULL DEFAULT '',
    year INTEGER NOT NULL DEFAULT 0,
    version TEXT NOT NULL DEFAULT '',
    engine TEXT NOT NULL DEFAULT '',
    transmission TEXT NOT NULL DEFAULT '',
    fuel TEXT NOT NULL DEFAULT '',
    km INTEGER NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE IF EXISTS vehicles;
