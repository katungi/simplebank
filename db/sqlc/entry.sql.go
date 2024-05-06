// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: entry.sql

package db

import (
	"context"
)

const createEntries = `-- name: CreateEntries :one
INSERT INTO entries (
    account_id, 
    amount
) VALUES (
    $1,$2
) RETURNING id, account_id, amount, created_at
`

type CreateEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntries, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntries = `-- name: DeleteEntries :exec
DELETE FROM entries
WHERE id = $1
`

func (q *Queries) DeleteEntries(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntries, id)
	return err
}

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, account_id, amount, created_at FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2 OFFSET $3
`

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntries = `-- name: UpdateEntries :exec
UPDATE entries
SET amount = $2
WHERE id = $1
RETURNING id, account_id, amount, created_at
`

type UpdateEntriesParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) UpdateEntries(ctx context.Context, arg UpdateEntriesParams) error {
	_, err := q.db.ExecContext(ctx, updateEntries, arg.ID, arg.Amount)
	return err
}
