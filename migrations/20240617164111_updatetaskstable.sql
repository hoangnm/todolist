-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks ADD COLUMN status TEXT DEFAULT 'todo' CHECK(status IN ('todo', 'inprogress', 'done'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tasks DROP COLUMN status;
-- +goose StatementEnd
