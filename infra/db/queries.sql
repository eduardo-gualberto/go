-- name: ListUsers :many
SELECT * FROM users;

-- name: ListParticipants :many
SELECT * FROM participants;

-- name: GetVersion :one
SELECT VERSION() AS version;