-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users(
    user_number, user_name
    ) VALUES (
        $1, $2
    ) RETURNING *;

-- name: ListParticipants :many
SELECT * FROM participants;

-- name: GetParticipantByID :one
SELECT * FROM participants
WHERE id = $1;

-- name: CreateParticipant :one
INSERT INTO participants(
    participant_number, participant_name, user_id
    ) VALUES (
        $1, $2, $3
    ) RETURNING *;

-- name: GetVersion :one
SELECT VERSION() AS version;