--1. Удалить БД
--DROP DATABASE DB_Alex;
--2. Создать БД
--CREATE DATABASE DB_Alex;
--3. Удалить Схему
--DROP SCHEMA schema CASCADE ;
-- 4. Создать схему
CREATE SCHEMA schema;
--5. Создать таблицу Users
CREATE TABLE schema.Users (
    id_user_main SERIAL PRIMARY KEY,
    name_user TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
--6. Создать таблицу Orders
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
--7. Создать таблицу Products
CREATE TABLE schema.Products (
    id_product_main SERIAL PRIMARY KEY,
    name_product TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL
);
--8. Создать таблицу Orders_Products
CREATE TABLE schema.Order_Products (
    id_order_f INT NOT NULL,
    id_product_f INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    PRIMARY KEY (id_order_f, id_product_f),
    CONSTRAINT fk_order FOREIGN KEY (id_order_f) REFERENCES schema.orders(id_order_main) ON DELETE CASCADE,
    CONSTRAINT fk_product FOREIGN KEY (id_product_f) REFERENCES schema.products(id_product_main) ON DELETE CASCADE
);
--9. Заполнить таблицу Users
INSERT INTO schema.Users (name_user, email, password)
VALUES
    ('Иван Иванов', 'ivan@example.com', 'qwerty'),
    ('Петя Смирнов', 'petr@example.com', '123'),
    ('Женя Жбанов', 'jbanov@example.com', '123'),
    ('Катя Жукова', 'jukova@example.com', '456'),
    ('Аня Гавриленко', 'gavrilenko@example.com', '456');
--10. Заполнить таблицу Products
INSERT INTO schema.Products (name_product, price)
VALUES
    ('Носки', 10.00),
    ('Ботинки', 120.00),
    ('Калоши', 50.00),
    ('Сандали', 70.00),
    ('Туфли',90.00);
-- 11. Заполнить таблицу Orders
INSERT INTO schema.Orders (id_user_f, order_date, total_amount)
VALUES
    (1, '2025-02-25', 0), -- Иван Иванов
    (2, '2025-02-25', 0), -- Петя Смирнов
    (3, '2025-02-25', 0), -- Женя Жбанов
    (4, '2025-02-25', 0), -- Катя Жукова
    (5, '2025-02-25', 0); -- Аня Гавриленко
-- 12. Заполнить таблицу Order_Products
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
-- 13. Удалить заказы с ID = 1
DELETE FROM schema.Orders WHERE id_order_main = 1;

-- 14. Установить цену 20.00, где ID =1 в таблице Products
UPDATE schema.Products
SET price = 20.00
WHERE id_product_main = 1;
-- 15. Обновить таблицу Users, установить фио для email jukova@example.com на Катю Муравьеву
UPDATE schema.Users
SET name_user = 'Екатерина Муравьева'
WHERE email = 'jukova@example.com';
-- 16. Удалить пользователя, где почта Жбанов
DELETE FROM schema.Users
WHERE email = 'jbanov@example.com';
-- 17. Удалить продукт, где цена менее 50
DELETE FROM schema.Products
WHERE price < 50;
-- 18. Обновить таблицу с заказами
UPDATE schema.Orders
SET total_amount = (
    SELECT COALESCE(SUM(op.quantity * p.price), 0)
    FROM schema.Order_Products op
             JOIN schema.Products p ON op.id_product_f = p.id_product_main
    WHERE op.id_order_f = Orders.id_order_main
);
-- 19. Выбрать пользователей, где пароль 123
SELECT * FROM schema.Users WHERE password LIKE '123';
-- 20. Выбрать продукты, где цена между 50 и 100
SELECT * FROM schema.Products WHERE price BETWEEN 50 AND 100;
-- 21. Показать ID заказа, дату заказа и общую сумму заказа (где ID = 4)
SELECT o.id_order_main, o.order_date, o.total_amount
FROM schema.Orders o
WHERE o.id_user_f = 4;
-- 22. Посчитать общие траты и среднюю цену товара для каждого пользователя
SELECT
    u.name_user,
    COALESCE(SUM(op.quantity * p.price), 0) AS total_spent,
    COALESCE(SUM(p.price * op.quantity) / NULLIF(SUM(op.quantity), 0), 0) AS avg_product_price
FROM schema.Users u
         LEFT JOIN schema.Orders o ON u.id_user_main = o.id_user_f
         LEFT JOIN schema.Order_Products op ON o.id_order_main = op.id_order_f
         LEFT JOIN schema.Products p ON op.id_product_f = p.id_product_main
GROUP BY u.name_user;
-- индексы
CREATE INDEX idx_users_name_user ON schema.Users(name_user);
CREATE INDEX idx_products_price ON schema.Products(price);
CREATE INDEX idx_order_products_order ON schema.Order_Products(id_order_f);
CREATE INDEX idx_orders_user ON schema.Orders(id_user_f);