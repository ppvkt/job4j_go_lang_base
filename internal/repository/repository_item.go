package repository

import (
	"context"
	"fmt"

	"job4j.ru/go-lang-base/internal/tracker"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepoPg struct {
	pool *pgxpool.Pool
}

func NewRepoPg(pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{pool: pool}
}

func (r *RepoPg) Create(ctx context.Context, it tracker.Item) error {
	_, err := r.pool.Exec(
		ctx,
		`insert into items(id, name) values($1, $2)`,
		it.ID, it.Name,
	)

	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}

	return nil
}

func (r *RepoPg) List(ctx context.Context) ([]tracker.Item, error) {
	rows, err := r.pool.Query(ctx, `select id, name from items`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []tracker.Item
	for rows.Next() {
		var item tracker.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RepoPg) Get(ctx context.Context, id string) (tracker.Item, error) {
	var it tracker.Item
	err := r.pool.QueryRow(
		ctx,
		"select id, name from items where id = $1",
		id,
	).Scan(&it.ID, &it.Name)

	return it, err
}

func (r *RepoPg) Delete(ctx context.Context, uuid string) error {
	_, err := r.pool.Exec(
		ctx,
		`DELETE FROM items WHERE id=$1`,
		uuid,
	)

	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}
	return nil
}

func (r *RepoPg) Update(ctx context.Context, uuid string, newName string) error {
	_, err := r.pool.Exec(
		ctx,
		`UPDATE items SET name=$1 WHERE id=$2`,
		newName,
		uuid,
	)

	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}
	return nil
}

func (r *RepoPg) Find(ctx context.Context, namePart string) []tracker.Item {
	pattern := "%" + namePart + "%"

	rows, err := r.pool.Query(
		ctx,
		`SELECT id, name FROM items  WHERE name ILIKE $1`,
		pattern,
	)

	if err != nil {
		return []tracker.Item{}
	}

	defer rows.Close()

	var items []tracker.Item
	for rows.Next() {
		var it tracker.Item
		if err := rows.Scan(&it.ID, &it.Name); err != nil {
			return []tracker.Item{}
		}
		items = append(items, it)
	}

	return items
}
