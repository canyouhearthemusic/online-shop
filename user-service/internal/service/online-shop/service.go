package onlineshop

import "github.com/canyouhearthemusic/online-shop/user-service/internal/domain/user"

type Service struct {
	repo user.Repository
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

func WithUserRepository(repo user.Repository) Configuration {
	return func(s *Service) error {
		s.repo = repo

		return nil
	}
}
