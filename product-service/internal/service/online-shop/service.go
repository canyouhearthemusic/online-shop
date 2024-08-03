package onlineshop

import "github.com/canyouhearthemusic/online-shop/product-service/internal/domain/product"

type Service struct {
	repo product.Repository
}

type Configuration func(s *Service) error

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

func WithProductRepository(repo product.Repository) Configuration {
	return func(s *Service) error {
		s.repo = repo

		return nil
	}
}
