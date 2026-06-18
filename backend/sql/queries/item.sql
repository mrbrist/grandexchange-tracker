-- name: AddItemHistory :exec
INSERT INTO ge_price_history (
        item_id,
        price_timestamp,
        avg_high_price,
        avg_low_price,
        high_volume,
        low_volume
    )
VALUES ($1, $2, $3, $4, $5, $6);
-- name: GetItemHistory :many
SELECT *
FROM ge_price_history
WHERE item_id = $1
ORDER BY price_timestamp ASC;
-- name: GetLatestTimestamp :one
SELECT COALESCE(MAX(price_timestamp), 0)::BIGINT AS latest_timestamp
FROM ge_price_history;
-- name: CountItemUpdates :one
SELECT COUNT(*) AS update_count
FROM ge_price_history
WHERE item_id = $1;
-- name: CountAllRecords :one
SELECT COUNT(*) AS total_records
FROM ge_price_history;
-- name: GetUniqueTimestamps :many
SELECT DISTINCT price_timestamp
FROM ge_price_history
ORDER BY price_timestamp;