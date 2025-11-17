-- name: PostBook :exec
INSERT INTO books (id, title, author, lang, category, bought_by, read_by) VALUES($1, $2, $3, $4, $5, $6, $7);