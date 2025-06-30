CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

INSERT INTO users (name,email,password)
VALUES
('naldy','naldy@mail.com','1234'),
('davinda','davinda@mail.com','1234'),
('ari','ari@mail.com','1234');
