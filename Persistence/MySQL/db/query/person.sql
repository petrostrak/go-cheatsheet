-- name: CreatePerson :execresult
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
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetPersonById :one
SELECT persons_name FROM person
WHERE person.id = ? LIMIT 1;

-- name: ListPersons :many
SELECT * FROM person
ORDER BY id;

-- name: UpdatePerson :execresult
UPDATE person
SET kind = ?, persons_name = ?, origins = ?, programming_languages = ?, tools = ?, 
linkedin = ?, github = ?, personal = ?, foreign_languages = ?, fav_food = ?, 
fav_drink = ?, fav_programming_language = ?, thinking_about = ?, hobbies = ?
WHERE person.id = ?; 

-- name: DeletePersonById :exec
DELETE FROM person
WHERE person.id = ?;