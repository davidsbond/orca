BEGIN;

CREATE TABLE IF NOT EXISTS task_run
(
    id              UUID      NOT NULL PRIMARY KEY,
    workflow_run_id UUID      NOT NULL REFERENCES workflow_run (id) ON DELETE CASCADE,
    task_name       TEXT      NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT NOW(),
    scheduled_at    TIMESTAMP,
    started_at      TIMESTAMP,
    completed_at    TIMESTAMP,
    status          NUMERIC   NOT NULL DEFAULT 0,
    input           JSONB     NOT NULL DEFAULT '{}',
    output          JSONB     NOT NULL DEFAULT '{}'
);

COMMIT;
