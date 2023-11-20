CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS global_settings (
  application_name VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS counter (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  adjustment_value INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE VIEW counter_view AS
SELECT
  c1.ID,
  c1.adjustment_value,
  SUM(c2.adjustment_value) AS current_value,
  c1.created_at
FROM
counter c1
JOIN
counter c2 ON c1.ID >= c2.ID
GROUP BY
c1.ID, c1.adjustment_value, c1.created_at;

INSERT INTO global_settings (application_name)
VALUES ('My Site');

INSERT INTO counter (adjustment_value)
VALUES (0);
