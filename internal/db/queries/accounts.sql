-- name: CreateAccount :exec
INSERT INTO accounts (id, email)
VALUES (?, ?);
