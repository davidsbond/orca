BEGIN;

CREATE TABLE IF NOT EXISTS worker
(
    id                UUID   NOT NULL PRIMARY KEY,
    advertise_address TEXT   NOT NULL UNIQUE,
    workflows         TEXT[] NOT NULL,
    tasks             TEXT[] NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_worker_workflows ON worker USING GIN (workflows);
CREATE INDEX IF NOT EXISTS idx_worker_tasks ON worker USING GIN (tasks);

COMMIT;
