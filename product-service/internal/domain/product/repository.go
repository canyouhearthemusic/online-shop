package product

import "context"

type Repository interface {
	List(ctx context.Context) ([]*Entity, error)
	GetByID(ctx context.Context, id string) (*Entity, error)
	Create(ctx context.Context, product *Entity) error
	Update(ctx context.Context, id string, product *Entity) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, product *Entity) ([]*Entity, error)
}
