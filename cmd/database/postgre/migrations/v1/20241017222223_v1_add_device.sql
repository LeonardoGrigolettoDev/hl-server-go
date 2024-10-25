-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE device (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    device_type VARCHAR(255) NOT NULL,
    mac VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE device;
-- +goose StatementEnd
