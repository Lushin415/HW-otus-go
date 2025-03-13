-- name: GetUsers :many
SELECT id_user_main, name_user, email FROM schema.Users;

-- name: GetUserByEmail :one
SELECT id_user_main, name_user, email, password
FROM schema.Users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO schema.Users (name_user, email, password)
VALUES ($1, $2, $3)
    RETURNING id_user_main;

-- name: GetOrdersByUser :many
SELECT id_order_main, order_date, total_amount
FROM schema.Orders
WHERE id_user_f = $1;

-- name: UpdateProductPrice :exec
UPDATE schema.Products
SET price = $1
WHERE id_product_main = $2;

-- name: DeleteUser :exec
DELETE FROM schema.Users WHERE id_user_main = $1;

-- name: GetUserByID :one
SELECT id_user_main, name_user, email FROM schema.Users WHERE id_user_main = $1;
