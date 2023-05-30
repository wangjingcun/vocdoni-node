// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: votes.sql

package indexerdb

import (
	"context"
	"database/sql"
	"time"

	"go.vocdoni.io/dvote/types"
	"go.vocdoni.io/dvote/vochain/state"
)

const createVote = `-- name: CreateVote :execresult
REPLACE INTO votes (
	nullifier, process_id, block_height, block_index,
	weight, voter_id, overwrite_count, creation_time
) VALUES (
	?, ?, ?, ?,
	?, ?, ?, ?
)
`

type CreateVoteParams struct {
	Nullifier      types.Nullifier
	ProcessID      types.ProcessID
	BlockHeight    int64
	BlockIndex     int64
	Weight         string
	VoterID        state.VoterID
	OverwriteCount int64
	CreationTime   time.Time
}

func (q *Queries) CreateVote(ctx context.Context, arg CreateVoteParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createVote,
		arg.Nullifier,
		arg.ProcessID,
		arg.BlockHeight,
		arg.BlockIndex,
		arg.Weight,
		arg.VoterID,
		arg.OverwriteCount,
		arg.CreationTime,
	)
}

const getVote = `-- name: GetVote :one
SELECT votes.nullifier, votes.process_id, votes.block_height, votes.block_index, votes.weight, votes.creation_time, votes.voter_id, votes.overwrite_count, transactions.hash FROM votes
LEFT JOIN transactions
	ON votes.block_height = transactions.block_height
	AND votes.block_index = transactions.block_index
WHERE nullifier = ?
LIMIT 1
`

type GetVoteRow struct {
	Nullifier      types.Nullifier
	ProcessID      types.ProcessID
	BlockHeight    int64
	BlockIndex     int64
	Weight         string
	CreationTime   time.Time
	VoterID        state.VoterID
	OverwriteCount int64
	Hash           types.Hash
}

func (q *Queries) GetVote(ctx context.Context, nullifier types.Nullifier) (GetVoteRow, error) {
	row := q.db.QueryRowContext(ctx, getVote, nullifier)
	var i GetVoteRow
	err := row.Scan(
		&i.Nullifier,
		&i.ProcessID,
		&i.BlockHeight,
		&i.BlockIndex,
		&i.Weight,
		&i.CreationTime,
		&i.VoterID,
		&i.OverwriteCount,
		&i.Hash,
	)
	return i, err
}

const searchVotes = `-- name: SearchVotes :many
SELECT votes.nullifier, votes.process_id, votes.block_height, votes.block_index, votes.weight, votes.creation_time, votes.voter_id, votes.overwrite_count, transactions.hash FROM votes
LEFT JOIN transactions
	ON votes.block_height = transactions.block_height
	AND votes.block_index = transactions.block_index
WHERE (? = '' OR process_id = ?)
	AND (? = '' OR (INSTR(LOWER(HEX(nullifier)), ?) > 0))
ORDER BY votes.block_height DESC, votes.nullifier ASC
LIMIT ?
OFFSET ?
`

type SearchVotesParams struct {
	ProcessID       types.ProcessID
	NullifierSubstr string
	Limit           int32
	Offset          int32
}

type SearchVotesRow struct {
	Nullifier      types.Nullifier
	ProcessID      types.ProcessID
	BlockHeight    int64
	BlockIndex     int64
	Weight         string
	CreationTime   time.Time
	VoterID        state.VoterID
	OverwriteCount int64
	Hash           types.Hash
}

func (q *Queries) SearchVotes(ctx context.Context, arg SearchVotesParams) ([]SearchVotesRow, error) {
	rows, err := q.db.QueryContext(ctx, searchVotes,
		arg.ProcessID,
		arg.ProcessID,
		arg.NullifierSubstr,
		arg.NullifierSubstr,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchVotesRow
	for rows.Next() {
		var i SearchVotesRow
		if err := rows.Scan(
			&i.Nullifier,
			&i.ProcessID,
			&i.BlockHeight,
			&i.BlockIndex,
			&i.Weight,
			&i.CreationTime,
			&i.VoterID,
			&i.OverwriteCount,
			&i.Hash,
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