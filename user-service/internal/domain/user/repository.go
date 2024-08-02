package user

import (
	"context"
)

type Repository interface {
	List(ctx context.Context) ([]*Entity, error)
	GetByID(ctx context.Context, id string) (*Entity, error)
	Create(ctx context.Context, user *Request) error
	Update(ctx context.Context, id string, user *Request) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, user *Request) ([]*Entity, error)
}
