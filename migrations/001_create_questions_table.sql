-- +goose Up
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE questions;