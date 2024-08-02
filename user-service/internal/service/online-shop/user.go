package onlineshop

import (
	"context"
	"fmt"

	"github.com/canyouhearthemusic/online-shop/user-service/internal/domain/user"
)

func (s *Service) ListUsers(ctx context.Context) (res []*user.Response, err error) {
	data, err := s.repo.List(ctx)
	if err != nil {
		fmt.Printf("failed to select: %v\n", err)
		return
	}

	res = user.ParseFromEntities(data)

	return
}

func (s *Service) CreateUser(ctx context.Context, req *user.Request) (err error) {
	err = s.repo.Create(ctx, req)
	if err != nil {
		err = fmt.Errorf("failed to create user: %v", err.Error())
	}

	return
}

func (s *Service) GetUser(ctx context.Context, id string) (res *user.Response, err error) {
	data, err := s.repo.GetByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("failed to get user: %v", err.Error())
	}

	res = user.ParseFromEntity(data)

	return
}

func (s *Service) UpdateUser(ctx context.Context, id string, req *user.Request) (err error) {
	err = s.repo.Update(ctx, id, req)
	if err != nil {
		err = fmt.Errorf("failed to update user: %v", err.Error())
	}

	return
}

func (s *Service) DeleteUser(ctx context.Context, id string) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		err = fmt.Errorf("failed to delete user: %v", err.Error())
	}

	return
}

func (s *Service) SearchUsers(ctx context.Context, req *user.Request) (res []*user.Response, err error) {
	data, err := s.repo.Search(ctx, req)
	if err != nil {
		err = fmt.Errorf("failed to search user: %v", err.Error())
		return
	}

	res = user.ParseFromEntities(data)

	return
}
