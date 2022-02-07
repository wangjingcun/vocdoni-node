// Code generated by sqlc. DO NOT EDIT.
// source: process.sql

package scrutinizerdb

import (
	"context"
	"database/sql"
	"time"

	"go.vocdoni.io/dvote/types"
)

const createProcess = `-- name: CreateProcess :execresult
INSERT INTO processes (
	id, entity_id, entity_index, start_block, end_block,
	results_height, have_results, final_results,
	census_root, rolling_census_root, rolling_census_size,
	max_census_size, census_uri, metadata,
	census_origin, status, namespace,
	envelope_pb, mode_pb, vote_opts_pb,
	private_keys, public_keys,
	question_index, creation_time,
	source_block_height, source_network_id
) VALUES (
	?, ?, ?, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?,
	?, ?,
	?, ?
)
`

type CreateProcessParams struct {
	ID                types.ProcessID
	EntityID          string
	EntityIndex       int64
	StartBlock        int64
	EndBlock          int64
	ResultsHeight     int64
	HaveResults       bool
	FinalResults      bool
	CensusRoot        types.CensusRoot
	RollingCensusRoot types.CensusRoot
	RollingCensusSize int64
	MaxCensusSize     int64
	CensusUri         string
	Metadata          string
	CensusOrigin      int64
	Status            int64
	Namespace         int64
	EnvelopePb        types.EncodedProtoBuf
	ModePb            types.EncodedProtoBuf
	VoteOptsPb        types.EncodedProtoBuf
	PrivateKeys       string
	PublicKeys        string
	QuestionIndex     int64
	CreationTime      time.Time
	SourceBlockHeight int64
	SourceNetworkID   string
}

func (q *Queries) CreateProcess(ctx context.Context, arg CreateProcessParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProcess,
		arg.ID,
		arg.EntityID,
		arg.EntityIndex,
		arg.StartBlock,
		arg.EndBlock,
		arg.ResultsHeight,
		arg.HaveResults,
		arg.FinalResults,
		arg.CensusRoot,
		arg.RollingCensusRoot,
		arg.RollingCensusSize,
		arg.MaxCensusSize,
		arg.CensusUri,
		arg.Metadata,
		arg.CensusOrigin,
		arg.Status,
		arg.Namespace,
		arg.EnvelopePb,
		arg.ModePb,
		arg.VoteOptsPb,
		arg.PrivateKeys,
		arg.PublicKeys,
		arg.QuestionIndex,
		arg.CreationTime,
		arg.SourceBlockHeight,
		arg.SourceNetworkID,
	)
}

const getProcess = `-- name: GetProcess :one
SELECT id, entity_id, entity_index, start_block, end_block, results_height, have_results, final_results, census_root, rolling_census_root, rolling_census_size, max_census_size, census_uri, metadata, census_origin, status, namespace, envelope_pb, mode_pb, vote_opts_pb, private_keys, public_keys, question_index, creation_time, source_block_height, source_network_id FROM processes
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetProcess(ctx context.Context, id types.ProcessID) (Process, error) {
	row := q.db.QueryRowContext(ctx, getProcess, id)
	var i Process
	err := row.Scan(
		&i.ID,
		&i.EntityID,
		&i.EntityIndex,
		&i.StartBlock,
		&i.EndBlock,
		&i.ResultsHeight,
		&i.HaveResults,
		&i.FinalResults,
		&i.CensusRoot,
		&i.RollingCensusRoot,
		&i.RollingCensusSize,
		&i.MaxCensusSize,
		&i.CensusUri,
		&i.Metadata,
		&i.CensusOrigin,
		&i.Status,
		&i.Namespace,
		&i.EnvelopePb,
		&i.ModePb,
		&i.VoteOptsPb,
		&i.PrivateKeys,
		&i.PublicKeys,
		&i.QuestionIndex,
		&i.CreationTime,
		&i.SourceBlockHeight,
		&i.SourceNetworkID,
	)
	return i, err
}

const getProcessStatus = `-- name: GetProcessStatus :one
SELECT status FROM processes
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetProcessStatus(ctx context.Context, id types.ProcessID) (int64, error) {
	row := q.db.QueryRowContext(ctx, getProcessStatus, id)
	var status int64
	err := row.Scan(&status)
	return status, err
}

const searchProcesses = `-- name: SearchProcesses :many
SELECT ID FROM processes
WHERE (LENGTH(?) = 0 OR entity_id = ?)
	AND (? = 0 OR namespace = ?)
	AND (? = 0 OR status = ?)
	AND (? = "" OR source_network_id = ?)
	-- TODO(mvdan): consider keeping an id_hex column for faster searches
	AND (? = "" OR (INSTR(LOWER(HEX(id)), ?) > 0))
	AND (? = FALSE OR have_results)
ORDER BY creation_time ASC, ID ASC
	-- TODO(mvdan): use sqlc.arg once limit/offset support it:
	-- https://github.com/kyleconroy/sqlc/issues/1025
LIMIT ?
OFFSET ?
`

type SearchProcessesParams struct {
	EntityID        string
	Namespace       int64
	Status          int64
	SourceNetworkID string
	IDSubstr        string
	WithResults     interface{}
	Limit           int32
	Offset          int32
}

func (q *Queries) SearchProcesses(ctx context.Context, arg SearchProcessesParams) ([]types.ProcessID, error) {
	rows, err := q.db.QueryContext(ctx, searchProcesses,
		arg.EntityID,
		arg.EntityID,
		arg.Namespace,
		arg.Namespace,
		arg.Status,
		arg.Status,
		arg.SourceNetworkID,
		arg.SourceNetworkID,
		arg.IDSubstr,
		arg.IDSubstr,
		arg.WithResults,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []types.ProcessID
	for rows.Next() {
		var id types.ProcessID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setProcessResultsCancelled = `-- name: SetProcessResultsCancelled :execresult
UPDATE processes
SET have_results = FALSE, final_results = TRUE
WHERE id = ?
`

func (q *Queries) SetProcessResultsCancelled(ctx context.Context, id types.ProcessID) (sql.Result, error) {
	return q.db.ExecContext(ctx, setProcessResultsCancelled, id)
}

const setProcessResultsHeight = `-- name: SetProcessResultsHeight :execresult
UPDATE processes
SET results_height = ?
WHERE id = ?
`

type SetProcessResultsHeightParams struct {
	ResultsHeight int64
	ID            types.ProcessID
}

func (q *Queries) SetProcessResultsHeight(ctx context.Context, arg SetProcessResultsHeightParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, setProcessResultsHeight, arg.ResultsHeight, arg.ID)
}

const setProcessResultsReady = `-- name: SetProcessResultsReady :execresult
UPDATE processes
SET have_results = TRUE, final_results = TRUE
WHERE id = ?
`

func (q *Queries) SetProcessResultsReady(ctx context.Context, id types.ProcessID) (sql.Result, error) {
	return q.db.ExecContext(ctx, setProcessResultsReady, id)
}

const updateProcessFromState = `-- name: UpdateProcessFromState :execresult
UPDATE processes
SET end_block           = ?,
	census_root         = ?,
	rolling_census_root = ?,
	census_uri          = ?,
	private_keys        = ?,
	public_keys         = ?,
	metadata            = ?,
	rolling_census_size = ?,
	status              = ?
WHERE id = ?
`

type UpdateProcessFromStateParams struct {
	EndBlock          int64
	CensusRoot        types.CensusRoot
	RollingCensusRoot types.CensusRoot
	CensusUri         string
	PrivateKeys       string
	PublicKeys        string
	Metadata          string
	RollingCensusSize int64
	Status            int64
	ID                types.ProcessID
}

func (q *Queries) UpdateProcessFromState(ctx context.Context, arg UpdateProcessFromStateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProcessFromState,
		arg.EndBlock,
		arg.CensusRoot,
		arg.RollingCensusRoot,
		arg.CensusUri,
		arg.PrivateKeys,
		arg.PublicKeys,
		arg.Metadata,
		arg.RollingCensusSize,
		arg.Status,
		arg.ID,
	)
}