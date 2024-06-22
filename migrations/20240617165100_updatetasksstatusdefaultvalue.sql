-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks_new (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    status TEXT DEFAULT 'todo'
);

INSERT INTO tasks_new(id, title, status) SELECT id, title, status FROM tasks;

DROP TABLE tasks;

ALTER TABLE tasks_new RENAME TO tasks;
-- +goose StatementEnd
