-- name: CreatePerson :one
INSERT INTO person (
    id,
    kind,
    persons_name,
    origins,
    programming_languages,
    tools,
    linkedin,
    github,
    personal,
    foreign_languages,
    fav_food,
    fav_drink,
    fav_programming_language,
    thinking_about,
    hobbies
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING *;

-- name: ReadPerson :many
SELECT persons_name FROM "person"
WHERE person.id = $1
ORDER BY "person".persons_name;

-- name: UpdatePerson :one
UPDATE "person" 
SET kind = $2, persons_name = $3, origins = $4, programming_languages = $5, tools = $6, 
linkedin = $7, github = $8, personal = $9, foreign_languages = $10, fav_food = $11, 
fav_drink = $12, fav_programming_language = $13, thinking_about = $14, hobbies = $15
WHERE "person".id = $1
RETURNING *; 

-- name: DeletePersonById :exec
DELETE FROM "person"
WHERE "person".id = $1;