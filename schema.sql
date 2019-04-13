-- NOTE: This is a PostgreSQL schema.

CREATE TABLE mailing_list (
    email varchar(64) NOT NULL,
    timestamp timestamp DEFAULT NOW(),
    PRIMARY KEY (email)
);