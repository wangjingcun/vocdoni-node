// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: token_fees.sql

package indexerdb

import (
	"context"
	"database/sql"
	"time"
)

const createTokenFee = `-- name: CreateTokenFee :execresult
INSERT INTO token_fees (
	from_account, block_height, reference,
	cost, tx_type, spend_time
) VALUES (
	?, ?, ?,
	?, ?, ?
)
`

type CreateTokenFeeParams struct {
	FromAccount []byte
	BlockHeight int64
	Reference   string
	Cost        int64
	TxType      string
	SpendTime   time.Time
}

func (q *Queries) CreateTokenFee(ctx context.Context, arg CreateTokenFeeParams) (sql.Result, error) {
	return q.exec(ctx, q.createTokenFeeStmt, createTokenFee,
		arg.FromAccount,
		arg.BlockHeight,
		arg.Reference,
		arg.Cost,
		arg.TxType,
		arg.SpendTime,
	)
}

const getTokenFees = `-- name: GetTokenFees :many
SELECT id, block_height, from_account, reference, cost, tx_type, spend_time FROM token_fees
ORDER BY spend_time DESC
LIMIT ?2
OFFSET ?1
`

type GetTokenFeesParams struct {
	Offset int64
	Limit  int64
}

func (q *Queries) GetTokenFees(ctx context.Context, arg GetTokenFeesParams) ([]TokenFee, error) {
	rows, err := q.query(ctx, q.getTokenFeesStmt, getTokenFees, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokenFee
	for rows.Next() {
		var i TokenFee
		if err := rows.Scan(
			&i.ID,
			&i.BlockHeight,
			&i.FromAccount,
			&i.Reference,
			&i.Cost,
			&i.TxType,
			&i.SpendTime,
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

const getTokenFeesByFromAccount = `-- name: GetTokenFeesByFromAccount :many
;

SELECT id, block_height, from_account, reference, cost, tx_type, spend_time FROM token_fees
WHERE from_account = ?1
ORDER BY spend_time DESC
LIMIT ?3
OFFSET ?2
`

type GetTokenFeesByFromAccountParams struct {
	FromAccount []byte
	Offset      int64
	Limit       int64
}

func (q *Queries) GetTokenFeesByFromAccount(ctx context.Context, arg GetTokenFeesByFromAccountParams) ([]TokenFee, error) {
	rows, err := q.query(ctx, q.getTokenFeesByFromAccountStmt, getTokenFeesByFromAccount, arg.FromAccount, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokenFee
	for rows.Next() {
		var i TokenFee
		if err := rows.Scan(
			&i.ID,
			&i.BlockHeight,
			&i.FromAccount,
			&i.Reference,
			&i.Cost,
			&i.TxType,
			&i.SpendTime,
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

const getTokenFeesByReference = `-- name: GetTokenFeesByReference :many
;

SELECT id, block_height, from_account, reference, cost, tx_type, spend_time FROM token_fees
WHERE reference = ?1
ORDER BY spend_time DESC
LIMIT ?3
OFFSET ?2
`

type GetTokenFeesByReferenceParams struct {
	Reference string
	Offset    int64
	Limit     int64
}

func (q *Queries) GetTokenFeesByReference(ctx context.Context, arg GetTokenFeesByReferenceParams) ([]TokenFee, error) {
	rows, err := q.query(ctx, q.getTokenFeesByReferenceStmt, getTokenFeesByReference, arg.Reference, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokenFee
	for rows.Next() {
		var i TokenFee
		if err := rows.Scan(
			&i.ID,
			&i.BlockHeight,
			&i.FromAccount,
			&i.Reference,
			&i.Cost,
			&i.TxType,
			&i.SpendTime,
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

const getTokenFeesByTxType = `-- name: GetTokenFeesByTxType :many
;

SELECT id, block_height, from_account, reference, cost, tx_type, spend_time FROM token_fees
WHERE tx_type = ?1
ORDER BY spend_time DESC
LIMIT ?3
OFFSET ?2
`

type GetTokenFeesByTxTypeParams struct {
	TxType string
	Offset int64
	Limit  int64
}

func (q *Queries) GetTokenFeesByTxType(ctx context.Context, arg GetTokenFeesByTxTypeParams) ([]TokenFee, error) {
	rows, err := q.query(ctx, q.getTokenFeesByTxTypeStmt, getTokenFeesByTxType, arg.TxType, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TokenFee
	for rows.Next() {
		var i TokenFee
		if err := rows.Scan(
			&i.ID,
			&i.BlockHeight,
			&i.FromAccount,
			&i.Reference,
			&i.Cost,
			&i.TxType,
			&i.SpendTime,
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