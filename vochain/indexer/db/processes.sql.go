// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: processes.sql

package indexerdb

import (
	"context"
	"database/sql"
	"time"

	"go.vocdoni.io/dvote/types"
)

const createProcess = `-- name: CreateProcess :execresult
INSERT INTO processes (
	id, entity_id, start_block, end_block,
	results_height, have_results, final_results,
	census_root, rolling_census_root, rolling_census_size,
	max_census_size, census_uri, metadata,
	census_origin, status, namespace,
	envelope_pb, mode_pb, vote_opts_pb,
	private_keys, public_keys,
	question_index, creation_time,
	source_block_height, source_network_id,

	results_votes, results_weight, results_envelope_height,
	results_block_height
) VALUES (
	?, ?, ?, ?,
	0, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?,
	?, ?,
	?, ?,

	?, '0', 0,
	0
)
`

type CreateProcessParams struct {
	ID                types.ProcessID
	EntityID          types.EntityID
	StartBlock        int64
	EndBlock          int64
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
	SourceNetworkID   int64
	ResultsVotes      string
}

func (q *Queries) CreateProcess(ctx context.Context, arg CreateProcessParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProcess,
		arg.ID,
		arg.EntityID,
		arg.StartBlock,
		arg.EndBlock,
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
		arg.ResultsVotes,
	)
}

const getEntityCount = `-- name: GetEntityCount :one
SELECT COUNT(DISTINCT entity_id) FROM processes
`

func (q *Queries) GetEntityCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getEntityCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getEntityProcessCount = `-- name: GetEntityProcessCount :one
SELECT COUNT(*) FROM processes
WHERE entity_id = ?1
`

func (q *Queries) GetEntityProcessCount(ctx context.Context, entityID types.EntityID) (int64, error) {
	row := q.db.QueryRowContext(ctx, getEntityProcessCount, entityID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getProcess = `-- name: GetProcess :one
SELECT id, entity_id, start_block, end_block, results_height, have_results, final_results, results_votes, results_weight, results_envelope_height, results_block_height, census_root, rolling_census_root, rolling_census_size, max_census_size, census_uri, metadata, census_origin, status, namespace, envelope_pb, mode_pb, vote_opts_pb, private_keys, public_keys, question_index, creation_time, source_block_height, source_network_id FROM processes
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetProcess(ctx context.Context, id types.ProcessID) (Process, error) {
	row := q.db.QueryRowContext(ctx, getProcess, id)
	var i Process
	err := row.Scan(
		&i.ID,
		&i.EntityID,
		&i.StartBlock,
		&i.EndBlock,
		&i.ResultsHeight,
		&i.HaveResults,
		&i.FinalResults,
		&i.ResultsVotes,
		&i.ResultsWeight,
		&i.ResultsEnvelopeHeight,
		&i.ResultsBlockHeight,
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

const getProcessCount = `-- name: GetProcessCount :one
SELECT COUNT(*) FROM processes
`

func (q *Queries) GetProcessCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getProcessCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getProcessIDsByFinalResults = `-- name: GetProcessIDsByFinalResults :many
;

SELECT id FROM processes
WHERE final_results = ?
`

func (q *Queries) GetProcessIDsByFinalResults(ctx context.Context, finalResults bool) ([]types.ProcessID, error) {
	rows, err := q.db.QueryContext(ctx, getProcessIDsByFinalResults, finalResults)
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

const searchEntities = `-- name: SearchEntities :many
SELECT DISTINCT entity_id FROM processes
WHERE (?1 = '' OR (INSTR(LOWER(HEX(entity_id)), ?1) > 0))
ORDER BY creation_time DESC, id ASC
LIMIT ?3
OFFSET ?2
`

type SearchEntitiesParams struct {
	EntityIDSubstr interface{}
	Offset         int64
	Limit          int64
}

func (q *Queries) SearchEntities(ctx context.Context, arg SearchEntitiesParams) ([]types.EntityID, error) {
	rows, err := q.db.QueryContext(ctx, searchEntities, arg.EntityIDSubstr, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []types.EntityID
	for rows.Next() {
		var entity_id types.EntityID
		if err := rows.Scan(&entity_id); err != nil {
			return nil, err
		}
		items = append(items, entity_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchProcesses = `-- name: SearchProcesses :many
SELECT id FROM processes
WHERE (?1 = 0 OR entity_id = ?2)
	AND (?3 = 0 OR namespace = ?3)
	AND (?4 = 0 OR status = ?4)
	AND (?5 = 0 OR source_network_id = ?5)
	-- TODO(mvdan): consider keeping an id_hex column for faster searches
	AND (?6 = '' OR (INSTR(LOWER(HEX(id)), ?6) > 0))
	AND (?7 = FALSE OR have_results)
ORDER BY creation_time DESC, id ASC
LIMIT ?9
OFFSET ?8
`

type SearchProcessesParams struct {
	EntityIDLen     interface{}
	EntityID        types.EntityID
	Namespace       interface{}
	Status          interface{}
	SourceNetworkID interface{}
	IDSubstr        interface{}
	WithResults     interface{}
	Offset          int64
	Limit           int64
}

// TODO(mvdan): when sqlc's parser is better, and does not get confused with
// string types, use:
// WHERE (LENGTH(sqlc.arg(entity_id)) = 0 OR entity_id = sqlc.arg(entity_id))
func (q *Queries) SearchProcesses(ctx context.Context, arg SearchProcessesParams) ([]types.ProcessID, error) {
	rows, err := q.db.QueryContext(ctx, searchProcesses,
		arg.EntityIDLen,
		arg.EntityID,
		arg.Namespace,
		arg.Status,
		arg.SourceNetworkID,
		arg.IDSubstr,
		arg.WithResults,
		arg.Offset,
		arg.Limit,
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
WHERE id = ?1
`

func (q *Queries) SetProcessResultsCancelled(ctx context.Context, id types.ProcessID) (sql.Result, error) {
	return q.db.ExecContext(ctx, setProcessResultsCancelled, id)
}

const setProcessResultsReady = `-- name: SetProcessResultsReady :execresult
UPDATE processes
SET have_results = TRUE, final_results = TRUE,
	results_votes = ?1,
	results_weight = ?2,
	results_block_height = ?3
WHERE id = ?4
`

type SetProcessResultsReadyParams struct {
	Votes       string
	Weight      string
	BlockHeight int64
	ID          types.ProcessID
}

func (q *Queries) SetProcessResultsReady(ctx context.Context, arg SetProcessResultsReadyParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, setProcessResultsReady,
		arg.Votes,
		arg.Weight,
		arg.BlockHeight,
		arg.ID,
	)
}

const updateProcessFromState = `-- name: UpdateProcessFromState :execresult
;

UPDATE processes
SET end_block           = ?1,
	census_root         = ?2,
	rolling_census_root = ?3,
	census_uri          = ?4,
	private_keys        = ?5,
	public_keys         = ?6,
	metadata            = ?7,
	rolling_census_size = ?8,
	status              = ?9
WHERE id = ?10
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

const updateProcessResultByID = `-- name: UpdateProcessResultByID :execresult
UPDATE processes
SET results_votes  = ?1,
    results_weight = ?2,
    vote_opts_pb = ?3,
    envelope_pb = ?4
WHERE id = ?5
`

type UpdateProcessResultByIDParams struct {
	Votes      string
	Weight     string
	VoteOptsPb types.EncodedProtoBuf
	EnvelopePb types.EncodedProtoBuf
	ID         types.ProcessID
}

func (q *Queries) UpdateProcessResultByID(ctx context.Context, arg UpdateProcessResultByIDParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProcessResultByID,
		arg.Votes,
		arg.Weight,
		arg.VoteOptsPb,
		arg.EnvelopePb,
		arg.ID,
	)
}

const updateProcessResults = `-- name: UpdateProcessResults :execresult
UPDATE processes
SET results_votes = ?1,
	results_weight = ?2,
	results_block_height = ?3
WHERE id = ?4 AND final_results = FALSE
`

type UpdateProcessResultsParams struct {
	Votes       string
	Weight      string
	BlockHeight int64
	ID          types.ProcessID
}

func (q *Queries) UpdateProcessResults(ctx context.Context, arg UpdateProcessResultsParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProcessResults,
		arg.Votes,
		arg.Weight,
		arg.BlockHeight,
		arg.ID,
	)
}
