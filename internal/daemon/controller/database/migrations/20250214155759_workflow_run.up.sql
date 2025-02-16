BEGIN;

CREATE TABLE IF NOT EXISTS workflow_run
(
    id                     UUID      NOT NULL PRIMARY KEY,
    parent_workflow_run_id UUID REFERENCES workflow_run (id) ON DELETE CASCADE,
    workflow_name          TEXT      NOT NULL,
    created_at             TIMESTAMP NOT NULL DEFAULT NOW(),
    scheduled_at           TIMESTAMP,
    started_at             TIMESTAMP,
    completed_at           TIMESTAMP,
    status                 NUMERIC   NOT NULL DEFAULT 0,
    input                  JSONB,
    output                 JSONB
);

COMMIT;
