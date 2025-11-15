CREATE TABLE books(
    id VARCHAR(26) PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    language VARCHAR(128) NOT NULL,
    category VARCHAR(128) NOT NULL,
    bought_by VARCHAR(255),
    read_by VARCHAR(255)
);