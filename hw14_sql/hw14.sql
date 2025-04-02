CREATE DATABASE DB_Alex;
CREATE SCHEMA schema;
CREATE TABLE schema.Users (
    id_user_main SERIAL PRIMARY KEY,
    name_user TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE schema.Orders (
    id_order_main SERIAL PRIMARY KEY,
    id_user_f INT NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (id_user_f)
            REFERENCES schema.Users(id_user_main)
            ON DELETE CASCADE,
    order_date DATE NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL
);

CREATE TABLE schema.Products (
    id_product_main SERIAL PRIMARY KEY,
    name_product TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL
);

CREATE TABLE schema.Order_Products (
    id_order_f INT NOT NULL,
    id_product_f INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    PRIMARY KEY (id_order_f, id_product_f),
    CONSTRAINT fk_order FOREIGN KEY (id_order_f) REFERENCES schema.orders(id_order_main) ON DELETE CASCADE,
    CONSTRAINT fk_product FOREIGN KEY (id_product_f) REFERENCES schema.products(id_product_main) ON DELETE CASCADE
);

INSERT INTO schema.Users (name_user, email, password)
VALUES
    ('Иван Иванов', 'ivan@example.com', 'qwerty'),
    ('Петя Смирнов', 'petr@example.com', '123'),
    ('Женя Жбанов', 'jbanov@example.com', '123'),
    ('Катя Жукова', 'jukova@example.com', '456'),
    ('Аня Гавриленко', 'gavrilenko@example.com', '456');

INSERT INTO schema.Products (name_product, price)
VALUES
    ('Носки', 10.00),
    ('Ботинки', 120.00),
    ('Калоши', 50.00),
    ('Сандали', 70.00),
    ('Туфли',90.00);

INSERT INTO schema.Orders (id_user_f, order_date, total_amount)
VALUES
    (1, '2025-02-25', 0), -- Иван Иванов
    (2, '2025-02-25', 0), -- Петя Смирнов
    (3, '2025-02-25', 0), -- Женя Жбанов
    (4, '2025-02-25', 0), -- Катя Жукова
    (5, '2025-02-25', 0); -- Аня Гавриленко

INSERT INTO schema.Order_Products (id_order_f, id_product_f, quantity)
VALUES
    -- Иван Иванов
    (1, 1, 1), -- 1 Носки
    (1, 2, 2), -- 2 Ботинки

    -- Петя Смирнов
    (2, 4, 3), -- 3 Сандали
    (2, 5, 1), -- 1 Туфли

    -- Женя Жбанов
    (3, 1, 5), -- 5 Носки

    -- Катя Жукова
    (4, 5, 1), -- 1 Туфли

    -- Аня Гавриленко
    (5, 5, 3), -- 3 Туфли
    (5, 1, 1); -- 1 Носки

DELETE FROM schema.Orders WHERE id_order_main = 1;


UPDATE schema.Products
SET price = 20.00
WHERE id_product_main = 1;

UPDATE schema.Users
SET name_user = 'Екатерина Муравьева'
WHERE email = 'jukova@example.com';

DELETE FROM schema.Users
WHERE email = 'jbanov@example.com';

DELETE FROM schema.Products
WHERE price < 50;

UPDATE schema.Orders
SET total_amount = (
    SELECT COALESCE(SUM(op.quantity * p.price), 0)
    FROM schema.Order_Products op
             JOIN schema.Products p ON op.id_product_f = p.id_product_main
    WHERE op.id_order_f = Orders.id_order_main
);

SELECT * FROM schema.Users WHERE password LIKE '123';

SELECT * FROM schema.Products WHERE price BETWEEN 50 AND 100;

SELECT o.id_order_main, o.order_date, o.total_amount
FROM schema.Orders o
WHERE o.id_user_f = 4;

SELECT
    u.name_user,
    COALESCE(SUM(op.quantity * p.price), 0) AS total_spent,
    COALESCE(SUM(p.price * op.quantity) / NULLIF(SUM(op.quantity), 0), 0) AS avg_product_price
FROM schema.Users u
         LEFT JOIN schema.Orders o ON u.id_user_main = o.id_user_f
         LEFT JOIN schema.Order_Products op ON o.id_order_main = op.id_order_f
         LEFT JOIN schema.Products p ON op.id_product_f = p.id_product_main
GROUP BY u.name_user;

CREATE INDEX idx_users_name_user ON schema.Users(name_user);
CREATE INDEX idx_products_price ON schema.Products(price);
CREATE INDEX idx_order_products_order ON schema.Order_Products(id_order_f);
CREATE INDEX idx_orders_user ON schema.Orders(id_user_f);