BEGIN;

-- We use a materialized view here as we don't need to constantly perform an UNNEST on the available tasks. We can
-- instead just repopulate this view whenever a worker joins/leaves.
CREATE MATERIALIZED VIEW IF NOT EXISTS task AS
SELECT DISTINCT name
FROM worker,
     UNNEST(tasks) AS name;

CREATE UNIQUE INDEX IF NOT EXISTS idx_task_name ON task (name);

-- Function that causes the materialized task view to refresh. Keeping it up-to-date with the latest state
-- of the workers registered with the controllers.
CREATE OR REPLACE FUNCTION refresh_task_view()
    RETURNS TRIGGER AS
$$
BEGIN
    REFRESH MATERIALIZED VIEW task;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Trigger that causes the view refresh whenever a worker joins or leaves.
CREATE OR REPLACE TRIGGER refresh_task_view
    AFTER INSERT OR UPDATE OR DELETE
    ON worker
    FOR EACH STATEMENT
EXECUTE FUNCTION refresh_task_view();

COMMIT;
