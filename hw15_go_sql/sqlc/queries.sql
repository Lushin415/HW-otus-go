-- name: CreateUser :one
INSERT INTO schema.Users (name_user, email, password)
VALUES ($1, $2, $3)
    RETURNING id_user_main;

-- name: CreateProduct :one
INSERT INTO schema.Products (name_product, price)
VALUES ($1, $2)
    RETURNING id_product_main;

-- name: CreateOrder :one
INSERT INTO schema.Orders (id_user_f, order_date, total_amount)
VALUES ($1, $2, $3)
    RETURNING id_order_main;

-- name: CreateOrderProduct :exec
INSERT INTO schema.Order_Products (id_order_f, id_product_f, quantity)
VALUES ($1, $2, $3);

-- name: DeleteOrder :exec
DELETE FROM schema.Orders WHERE id_order_main = $1;

-- name: UpdateProductPrice :exec
UPDATE schema.Products
SET price = $1
WHERE id_product_main = $2;

-- name: UpdateUserName :exec
UPDATE schema.Users
SET name_user = $1
WHERE email = $2;

-- name: DeleteUser :exec
DELETE FROM schema.Users WHERE email = $1;

-- name: DeleteCheapProducts :exec
DELETE FROM schema.Products WHERE price < $1;

-- name: UpdateOrderTotal :exec
UPDATE schema.Orders
SET total_amount = (
    SELECT COALESCE(SUM(op.quantity * p.price), 0)
    FROM schema.Order_Products op
             JOIN schema.Products p ON op.id_product_f = p.id_product_main
    WHERE op.id_order_f = schema.Orders.id_order_main
);

-- name: GetUsersByPassword :many
SELECT * FROM schema.Users WHERE password = $1;

-- name: GetProductsByPriceRange :many
SELECT * FROM schema.Products WHERE price BETWEEN $1 AND $2;

-- name: GetOrderByUserID :many
SELECT id_order_main, order_date, total_amount
FROM schema.Orders
WHERE id_user_f = $1;

-- name: GetUserSpendingStats :many
SELECT
    u.name_user,
    COALESCE(SUM(op.quantity * p.price), 0) AS total_spent,
    COALESCE(SUM(p.price * op.quantity) / NULLIF(SUM(op.quantity), 0), 0) AS avg_product_price
FROM schema.Users u
         LEFT JOIN schema.Orders o ON u.id_user_main = o.id_user_f
         LEFT JOIN schema.Order_Products op ON o.id_order_main = op.id_order_f
         LEFT JOIN schema.Products p ON op.id_product_f = p.id_product_main
GROUP BY u.name_user;
