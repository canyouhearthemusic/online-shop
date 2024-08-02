package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/canyouhearthemusic/online-shop/user-service/internal/domain/user"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) List(ctx context.Context) ([]*user.Entity, error) {
	var users []*user.Entity

	if err := r.db.SelectContext(ctx, &users, "SELECT id, name, email from users"); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*user.Entity, error) {
	var user user.Entity

	if err := r.db.GetContext(ctx, &user, "SELECT id, name, email from users WHERE id = $1", id); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *user.Request) error {
	query := `INSERT INTO users(id, name, email) VALUES ($1, $2, $3)`

	args := []any{uuid.NewString(), user.Name, user.Email}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, id string, user *user.Request) error {
	query := `UPDATE users SET name = $1, email = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3`

	args := []any{user.Name, user.Email, id}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	if _, err := r.db.ExecContext(ctx, "DELETE FROM users where id = $1", id); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Search(ctx context.Context, req *user.Request) ([]*user.Entity, error) {
	var users []*user.Entity

	query := "SELECT id, name, email FROM users WHERE 1=1"

	sets, args := r.prepareArgs(req)
	if len(sets) > 0 {
		query += " AND " + strings.Join(sets, " AND ")
	}

	if err := r.db.SelectContext(ctx, &users, query, args...); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) prepareArgs(user *user.Request) (sets []string, args []any) {
	if user.Name != "" {
		args = append(args, user.Name)
		sets = append(sets, fmt.Sprintf("name=$%d", len(args)))
	}

	if user.Email != "" {
		args = append(args, user.Email)
		sets = append(sets, fmt.Sprintf("email=$%d", len(args)))
	}

	return
}
