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

-- name: GetUsersByPasswordPattern :many
SELECT id_user_main, name_user, email
FROM schema.Users
WHERE password LIKE $1;

-- name: GetProductsByPriceRange :many
SELECT id_product_main, name_product, price
FROM schema.Products
WHERE price BETWEEN $1 AND $2;

-- name: GetUserSpendingMetrics :many
SELECT
    u.id_user_main,
    u.name_user,
    COALESCE(SUM(op.quantity * p.price), 0) AS total_spent,
    COALESCE(SUM(p.price * op.quantity) / NULLIF(SUM(op.quantity), 0), 0) AS avg_product_price
FROM schema.Users u
         LEFT JOIN schema.Orders o ON u.id_user_main = o.id_user_f
         LEFT JOIN schema.Order_Products op ON o.id_order_main = op.id_order_f
         LEFT JOIN schema.Products p ON op.id_product_f = p.id_product_main
GROUP BY u.id_user_main, u.name_user;