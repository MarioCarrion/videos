-- name: CreateUser :one
INSERT INTO users (
	name,
	birth_year
) VALUES (
	@name,
	@birth_year
)
RETURNING id;

-- name: GetUser :one
SELECT
	id,
	name,
	birth_year
FROM users
WHERE
	id = @id;
