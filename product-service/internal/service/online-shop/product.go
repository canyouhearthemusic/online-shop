package onlineshop

import (
	"context"
	"fmt"

	"github.com/canyouhearthemusic/online-shop/product-service/internal/domain/product"
)

func (s *Service) ListProducts(ctx context.Context) (res []*product.Entity, err error) {
	res, err = s.repo.List(ctx)
	if err != nil {
		err = fmt.Errorf("failed to select products: %v", err.Error())
	}

	return
}

func (s *Service) CreateProduct(ctx context.Context, req *product.Entity) (err error) {
	err = s.repo.Create(ctx, req)
	if err != nil {
		err = fmt.Errorf("failed to create product: %v", err.Error())
	}

	return
}

func (s *Service) GetProduct(ctx context.Context, id string) (res *product.Entity, err error) {
	res, err = s.repo.GetByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("failed to get product: %v", err.Error())
	}

	return
}

func (s *Service) UpdateProduct(ctx context.Context, id string, req *product.Entity) (err error) {
	err = s.repo.Update(ctx, id, req)
	if err != nil {
		err = fmt.Errorf("failed to update product: %v", err.Error())
	}

	return
}

func (s *Service) DeleteProduct(ctx context.Context, id string) (err error) {
	err = s.repo.Delete(ctx, id)
	if err != nil {
		err = fmt.Errorf("failed to delete product: %v", err.Error())
	}

	return
}

func (s *Service) SearchProducts(ctx context.Context, req *product.Entity) (res []*product.Entity, err error) {
	res, err = s.repo.Search(ctx, req)
	if err != nil {
		err = fmt.Errorf("failed to search product: %v", err.Error())
		return
	}

	return
}
