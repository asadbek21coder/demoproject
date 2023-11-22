CREATE TABLE IF NOT EXISTS "book" (
    book_id SERIAL NOT NULL PRIMARY KEY,
    book_title VARCHAR(255) NOT NULL,
    book_author INT NOT NULL REFERENCES "author" ("author_id"),
    book_price INT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITHOUT TIME ZONE

);