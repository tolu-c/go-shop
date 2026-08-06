-- name: ListProductss :many
SELECT * FROM products;
-- ORDER BY name;

-- name FindProductById :one
SELECT * FROM products WHERE id = $1;
