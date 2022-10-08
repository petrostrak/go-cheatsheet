package db

import (
	"context"
)

type Querier interface {
	CreatePerson(ctx context.Context, arg CreatePersonParams) (Person, error)
	DeletePersonById(ctx context.Context, id int64) error
	ReadPerson(ctx context.Context, id int64) ([]string, error)
	UpdatePerson(ctx context.Context, arg UpdatePersonParams) (Person, error)
}

var _ Querier = (*Queries)(nil)
