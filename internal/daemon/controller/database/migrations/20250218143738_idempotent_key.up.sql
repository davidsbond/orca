BEGIN;

ALTER TABLE IF EXISTS task_run
    ADD COLUMN IF NOT EXISTS idempotent_key TEXT;
ALTER TABLE IF EXISTS workflow_run
    ADD COLUMN IF NOT EXISTS idempotent_key TEXT;

CREATE INDEX IF NOT EXISTS idx_task_run_task_name ON task_run (task_name);
CREATE INDEX IF NOT EXISTS idx_task_run_idempotent_key ON task_run (idempotent_key);

CREATE INDEX IF NOT EXISTS idx_workflow_run_workflow_name ON workflow_run (workflow_name);
CREATE INDEX IF NOT EXISTS idx_workflow_run_idempotent_key ON workflow_run (idempotent_key);

COMMIT;
