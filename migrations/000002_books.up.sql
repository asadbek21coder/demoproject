CREATE TABLE IF NOT EXISTS "books" (
    id SERIAL NOT NULL primary key,
    name VARCHAR(128) NOT NULL,
    author VARCHAR(128) NOT NULL,
    price INT NOT NULL
)