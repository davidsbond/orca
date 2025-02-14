BEGIN;

CREATE TABLE IF NOT EXISTS workflow_run
(
    id            UUID      NOT NULL PRIMARY KEY,
    workflow_name TEXT      NOT NULL,
    scheduled_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    started_at    TIMESTAMP,
    completed_at  TIMESTAMP,
    status        NUMERIC   NOT NULL DEFAULT 0,
    input         JSONB     NOT NULL DEFAULT '{}',
    output        JSONB     NOT NULL DEFAULT '{}'
);

COMMIT;
