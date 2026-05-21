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
WHERE item_id = $1;