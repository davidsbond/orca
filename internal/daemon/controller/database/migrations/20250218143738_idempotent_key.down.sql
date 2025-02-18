BEGIN;

ALTER TABLE IF EXISTS task_run
    DROP COLUMN IF EXISTS idempotent_key;
ALTER TABLE IF EXISTS workflow_run
    DROP COLUMN IF EXISTS idempotent_key;

COMMIT;
