CREATE TABLE books(
    id VARCHAR(26) PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    author VARCHAR(128) NOT NULL,
    lang VARCHAR(128) NOT NULL,
    category VARCHAR(128) NOT NULL,
    bought_by VARCHAR(255) NOT NULL,
    read_by VARCHAR(255)
);