package repository

import (
	"github.com/canyouhearthemusic/online-shop/common/store"
	"github.com/canyouhearthemusic/online-shop/user-service/internal/domain/user"
	"github.com/canyouhearthemusic/online-shop/user-service/internal/repository/postgres"
)

type Repository struct {
	postgres store.SQLX

	User user.Repository
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

		// if err = store.Migrate(dsn); err != nil {
		// 	return
		// }

		r.User = postgres.NewUserRepository(r.postgres.Client)

		return
	}
}
