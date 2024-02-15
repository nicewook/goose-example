-- 사용자 데이터 삽입
INSERT INTO user (username, email, password_hash)
VALUES ('user1', 'user1@example.com', SHA2('password1', 256)),
       ('user2', 'user2@example.com', SHA2('password2', 256)),
       ('user3', 'user3@example.com', SHA2('password3', 256));

-- 제품 데이터 삽입
INSERT INTO product (name, description, price, stock_quantity)
VALUES ('Product 1', 'Description for product 1', 10.00, 50),
       ('Product 2', 'Description for product 2', 15.00, 100),
       ('Product 3', 'Description for product 3', 20.00, 150);
