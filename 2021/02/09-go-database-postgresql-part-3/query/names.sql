-- name: SelectName :one
SELECT
  nconst,
  primary_name,
  birth_year,
  death_year
FROM
  names
WHERE
  nconst = @nconst
LIMIT 1;
