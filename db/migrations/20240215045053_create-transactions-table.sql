-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transaction
(
    id               SERIAL PRIMARY KEY,
    user_id          INT            NOT NULL,
    amount           DECIMAL(10, 2) NOT NULL,
    transaction_type VARCHAR(50)    NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transaction;
-- +goose StatementEnd
