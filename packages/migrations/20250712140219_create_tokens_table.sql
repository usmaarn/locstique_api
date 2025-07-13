-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tokens(
    id TEXT NOT NULL PRIMARY KEY,
    user_id uuid NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    user_agent text,
    ip_address text,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tokens;
-- +goose StatementEnd
