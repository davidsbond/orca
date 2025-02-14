BEGIN;

CREATE TABLE IF NOT EXISTS worker
(
    id                UUID   NOT NULL PRIMARY KEY,
    advertise_address TEXT   NOT NULL,
    workflows         TEXT[] NOT NULL,
    tasks             TEXT[] NOT NULL
);

COMMIT;
