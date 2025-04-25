CREATE TABLE IF NOT EXISTS sample_table
(
    id integer
        constraint sample_table_pk
            primary key,
    name text
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_sample_table_name
ON sample_table ("name");
