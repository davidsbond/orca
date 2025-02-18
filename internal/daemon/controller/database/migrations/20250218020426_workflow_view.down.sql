BEGIN;

DROP TRIGGER IF EXISTS refresh_workflow_view ON worker;
DROP FUNCTION IF EXISTS refresh_workflow_view;
DROP MATERIALIZED VIEW IF EXISTS workflow;

COMMIT;
