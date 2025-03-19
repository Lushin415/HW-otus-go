--1. Удалить БД
--DROP DATABASE DB_Alex;
--2. Создать БД
CREATE DATABASE DB_Alex;
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