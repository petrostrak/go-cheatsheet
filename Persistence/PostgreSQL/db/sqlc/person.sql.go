package db

import (
	"context"

	"github.com/lib/pq"
)

const createPerson = `-- name: CreatePerson :one
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
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id, kind, persons_name, origins, programming_languages, tools, linkedin, github, personal, foreign_languages, fav_food, fav_drink, fav_programming_language, thinking_about, hobbies
`

type CreatePersonParams struct {
	ID                     int64    `json:"id"`
	Kind                   string   `json:"kind"`
	PersonsName            string   `json:"persons_name"`
	Origins                string   `json:"origins"`
	ProgrammingLanguages   []string `json:"programming_languages"`
	Tools                  []string `json:"tools"`
	Linkedin               string   `json:"linkedin"`
	Github                 string   `json:"github"`
	Personal               string   `json:"personal"`
	ForeignLanguages       []string `json:"foreign_languages"`
	FavFood                string   `json:"fav_food"`
	FavDrink               string   `json:"fav_drink"`
	FavProgrammingLanguage string   `json:"fav_programming_language"`
	ThinkingAbout          []string `json:"thinking_about"`
	Hobbies                []string `json:"hobbies"`
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, createPerson,
		arg.ID,
		arg.Kind,
		arg.PersonsName,
		arg.Origins,
		pq.Array(arg.ProgrammingLanguages),
		pq.Array(arg.Tools),
		arg.Linkedin,
		arg.Github,
		arg.Personal,
		pq.Array(arg.ForeignLanguages),
		arg.FavFood,
		arg.FavDrink,
		arg.FavProgrammingLanguage,
		pq.Array(arg.ThinkingAbout),
		pq.Array(arg.Hobbies),
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Kind,
		&i.PersonsName,
		&i.Origins,
		pq.Array(&i.ProgrammingLanguages),
		pq.Array(&i.Tools),
		&i.Linkedin,
		&i.Github,
		&i.Personal,
		pq.Array(&i.ForeignLanguages),
		&i.FavFood,
		&i.FavDrink,
		&i.FavProgrammingLanguage,
		pq.Array(&i.ThinkingAbout),
		pq.Array(&i.Hobbies),
	)
	return i, err
}

const deletePersonById = `-- name: DeletePersonById :exec
DELETE FROM "person"
WHERE "person".id = $1
`

func (q *Queries) DeletePersonById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePersonById, id)
	return err
}

const readPerson = `-- name: ReadPerson :many
SELECT persons_name FROM "person"
WHERE person.id = $1
ORDER BY "person".persons_name
`

func (q *Queries) ReadPerson(ctx context.Context, id int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, readPerson, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var persons_name string
		if err := rows.Scan(&persons_name); err != nil {
			return nil, err
		}
		items = append(items, persons_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePerson = `-- name: UpdatePerson :one
UPDATE "person" 
SET kind = $2, persons_name = $3, origins = $4, programming_languages = $5, tools = $6, 
linkedin = $7, github = $8, personal = $9, foreign_languages = $10, fav_food = $11, 
fav_drink = $12, fav_programming_language = $13, thinking_about = $14, hobbies = $15
WHERE "person".id = $1
RETURNING id, kind, persons_name, origins, programming_languages, tools, linkedin, github, personal, foreign_languages, fav_food, fav_drink, fav_programming_language, thinking_about, hobbies
`

type UpdatePersonParams struct {
	ID                     int64    `json:"id"`
	Kind                   string   `json:"kind"`
	PersonsName            string   `json:"persons_name"`
	Origins                string   `json:"origins"`
	ProgrammingLanguages   []string `json:"programming_languages"`
	Tools                  []string `json:"tools"`
	Linkedin               string   `json:"linkedin"`
	Github                 string   `json:"github"`
	Personal               string   `json:"personal"`
	ForeignLanguages       []string `json:"foreign_languages"`
	FavFood                string   `json:"fav_food"`
	FavDrink               string   `json:"fav_drink"`
	FavProgrammingLanguage string   `json:"fav_programming_language"`
	ThinkingAbout          []string `json:"thinking_about"`
	Hobbies                []string `json:"hobbies"`
}

func (q *Queries) UpdatePerson(ctx context.Context, arg UpdatePersonParams) (Person, error) {
	row := q.db.QueryRowContext(ctx, updatePerson,
		arg.ID,
		arg.Kind,
		arg.PersonsName,
		arg.Origins,
		pq.Array(arg.ProgrammingLanguages),
		pq.Array(arg.Tools),
		arg.Linkedin,
		arg.Github,
		arg.Personal,
		pq.Array(arg.ForeignLanguages),
		arg.FavFood,
		arg.FavDrink,
		arg.FavProgrammingLanguage,
		pq.Array(arg.ThinkingAbout),
		pq.Array(arg.Hobbies),
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Kind,
		&i.PersonsName,
		&i.Origins,
		pq.Array(&i.ProgrammingLanguages),
		pq.Array(&i.Tools),
		&i.Linkedin,
		&i.Github,
		&i.Personal,
		pq.Array(&i.ForeignLanguages),
		&i.FavFood,
		&i.FavDrink,
		&i.FavProgrammingLanguage,
		pq.Array(&i.ThinkingAbout),
		pq.Array(&i.Hobbies),
	)
	return i, err
}
