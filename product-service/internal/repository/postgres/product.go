package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/canyouhearthemusic/online-shop/product-service/internal/domain/product"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) List(ctx context.Context) ([]*product.Entity, error) {
	var products []*product.Entity
	query := "SELECT id, title, description, price, category, quantity FROM products"

	if err := r.db.SelectContext(ctx, &products, query); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*product.Entity, error) {
	var res product.Entity
	query := "SELECT id, title, description, price, category, quantity FROM products WHERE id = $1"

	if err := r.db.GetContext(ctx, &res, query, id); err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *ProductRepository) Create(ctx context.Context, req *product.Entity) (err error) {
	query := "INSERT INTO products(id, title, description, price, category, quantity) VALUES (:id, :title, :description, :price, :category, :quantity)"
	req.ID = uuid.NewString()

	if _, err = r.db.NamedExecContext(ctx, query, req); err != nil {
		return
	}

	return
}

func (r *ProductRepository) Update(ctx context.Context, id string, req *product.Entity) (err error) {
	query := "UPDATE products SET title = :title, description = :description, price = :price, category = :category, quantity = :quantity WHERE id = :id"

	if _, err = r.db.NamedExecContext(ctx, query, req); err != nil {
		return
	}

	return
}

func (r *ProductRepository) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM products WHERE id = $1"

	if _, err = r.db.ExecContext(ctx, query, id); err != nil {
		return
	}

	return
}

func (r *ProductRepository) Search(ctx context.Context, req *product.Entity) ([]*product.Entity, error) {
	var products []*product.Entity

	query := "SELECT id, title, description, price, category, quantity FROM products WHERE 1=1"

	sets, args := r.prepareArgs(req)
	if len(sets) > 0 {
		query += " AND " + strings.Join(sets, " AND ")
	}

	if err := r.db.SelectContext(ctx, &products, query, args...); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) prepareArgs(product *product.Entity) (sets []string, args []any) {
	if product.Title != "" {
		args = append(args, product.Title)
		sets = append(sets, fmt.Sprintf("title=$%d", len(args)))
	}

	if product.Category != "" {
		args = append(args, product.Category)
		sets = append(sets, fmt.Sprintf("category=$%d", len(args)))
	}

	return
}
