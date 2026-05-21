-- +goose Up
CREATE TABLE ge_price_history (
    id BIGSERIAL PRIMARY KEY,
    item_id INTEGER NOT NULL,
    price_timestamp TIMESTAMP NOT NULL,
    -- price data
    avg_high_price INTEGER NOT NULL,
    avg_low_price INTEGER NOT NULL,
    high_volume BIGINT NOT NULL,
    low_volume BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    -- prevent duplicate snapshots
    UNIQUE (item_id, price_timestamp)
);
-- +goose Down
DROP TABLE ge_price_history;