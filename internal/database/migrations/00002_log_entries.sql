-- +goose Up
CREATE TABLE log_entries (
    id SERIAL PRIMARY KEY,
    vehicle_id INTEGER NOT NULL REFERENCES vehicles(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    notes TEXT NOT NULL DEFAULT '',
    type TEXT NOT NULL,
    date DATE NOT NULL,
    km INTEGER NOT NULL,
    cost INTEGER NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS log_entries;
