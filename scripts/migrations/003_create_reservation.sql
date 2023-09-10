CREATE TABLE reservations (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    time_from   TIMESTAMP NOT NULL,
    time_to     TIMESTAMP NOT NULL,
    resv_by     INTEGER NOT NULL REFERENCES users(id),
    description TEXT,
    location    VARCHAR(255),
    color       VARCHAR(50),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
