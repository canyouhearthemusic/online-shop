package repository

import (
	"github.com/canyouhearthemusic/online-shop/common/store"
	"github.com/canyouhearthemusic/online-shop/product-service/internal/domain/product"
	"github.com/canyouhearthemusic/online-shop/product-service/internal/repository/postgres"
)

type Repository struct {
	postgres store.SQLX

	Product product.Repository
}

type Configuration func(r *Repository) error

func New(configs ...Configuration) (r *Repository, err error) {
	r = &Repository{}

	for _, cfg := range configs {
		if err = cfg(r); err != nil {
			return
		}
	}

	return
}

func (r *Repository) Close() {
	if r.postgres.Client != nil {
		r.postgres.Client.Close()
	}
}

func WithPostgresStore(dsn string) Configuration {
	return func(r *Repository) (err error) {
		r.postgres, err = store.NewSQL(dsn)
		if err != nil {
			return
		}

		r.Product = postgres.NewProductRepository(r.postgres.Client)

		return
	}
}
