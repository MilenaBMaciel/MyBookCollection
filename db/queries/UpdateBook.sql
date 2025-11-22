-- name: UpdateBook :exec
UPDATE books 
SET title=$2, author=$3, lang=$4, category=$5, bought_by=$6, read_by=$7
WHERE id=$1;