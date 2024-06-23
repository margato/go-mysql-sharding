CREATE DATABASE customers;
USE customers;

CREATE TABLE customers
(
    id         VARCHAR(26) NOT NULL, -- Study only. Do not use varchar in production, use Binary instead.
    name       VARCHAR(30) NOT NULL,
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
