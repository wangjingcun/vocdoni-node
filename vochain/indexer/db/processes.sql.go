// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: processes.sql

package indexerdb

import (
	"context"
	"database/sql"
	"time"

	"go.vocdoni.io/dvote/types"
)

const computeProcessVoteCount = `-- name: ComputeProcessVoteCount :execresult
UPDATE processes
SET vote_count = (SELECT COUNT(*) FROM votes WHERE process_id = id)
WHERE id = ?1
`

func (q *Queries) ComputeProcessVoteCount(ctx context.Context, id types.ProcessID) (sql.Result, error) {
	return q.exec(ctx, q.computeProcessVoteCountStmt, computeProcessVoteCount, id)
}

const createProcess = `-- name: CreateProcess :execresult
INSERT INTO processes (
	id, entity_id, start_block, end_block, start_date, end_date, 
	block_count, vote_count, have_results, final_results, census_root,
	max_census_size, census_uri, metadata,
	census_origin, status, namespace,
	envelope, mode, vote_opts,
	private_keys, public_keys,
	question_index, creation_time,
	source_block_height, source_network_id,
	from_archive, chain_id,

	results_votes, results_weight, results_block_height
) VALUES (
	?, ?, ?, ?, ?, ?,
	?, ?, ?, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?, ?,
	?, ?,
	?, ?,
	?, ?,
	?, ?,

	?, '"0"', 0
)
`

type CreateProcessParams struct {
	ID                types.ProcessID
	EntityID          types.EntityID
	StartBlock        int64
	EndBlock          int64
	StartDate         time.Time
	EndDate           time.Time
	BlockCount        int64
	VoteCount         int64
	HaveResults       bool
	FinalResults      bool
	CensusRoot        types.CensusRoot
	MaxCensusSize     int64
	CensusUri         string
	Metadata          string
	CensusOrigin      int64
	Status            int64
	Namespace         int64
	Envelope          string
	Mode              string
	VoteOpts          string
	PrivateKeys       string
	PublicKeys        string
	QuestionIndex     int64
	CreationTime      time.Time
	SourceBlockHeight int64
	SourceNetworkID   int64
	FromArchive       bool
	ChainID           string
	ResultsVotes      string
}

func (q *Queries) CreateProcess(ctx context.Context, arg CreateProcessParams) (sql.Result, error) {
	return q.exec(ctx, q.createProcessStmt, createProcess,
		arg.ID,
		arg.EntityID,
		arg.StartBlock,
		arg.EndBlock,
		arg.StartDate,
		arg.EndDate,
		arg.BlockCount,
		arg.VoteCount,
		arg.HaveResults,
		arg.FinalResults,
		arg.CensusRoot,
		arg.MaxCensusSize,
		arg.CensusUri,
		arg.Metadata,
		arg.CensusOrigin,
		arg.Status,
		arg.Namespace,
		arg.Envelope,
		arg.Mode,
		arg.VoteOpts,
		arg.PrivateKeys,
		arg.PublicKeys,
		arg.QuestionIndex,
		arg.CreationTime,
		arg.SourceBlockHeight,
		arg.SourceNetworkID,
		arg.FromArchive,
		arg.ChainID,
		arg.ResultsVotes,
	)
}

const getEntityCount = `-- name: GetEntityCount :one
SELECT COUNT(DISTINCT entity_id) FROM processes
`

func (q *Queries) GetEntityCount(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getEntityCountStmt, getEntityCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getProcess = `-- name: GetProcess :one
SELECT id, entity_id, start_block, end_block, start_date, end_date, block_count, vote_count, chain_id, have_results, final_results, results_votes, results_weight, results_block_height, census_root, max_census_size, census_uri, metadata, census_origin, status, namespace, envelope, mode, vote_opts, private_keys, public_keys, question_index, creation_time, source_block_height, source_network_id, from_archive FROM processes
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetProcess(ctx context.Context, id types.ProcessID) (Process, error) {
	row := q.queryRow(ctx, q.getProcessStmt, getProcess, id)
	var i Process
	err := row.Scan(
		&i.ID,
		&i.EntityID,
		&i.StartBlock,
		&i.EndBlock,
		&i.StartDate,
		&i.EndDate,
		&i.BlockCount,
		&i.VoteCount,
		&i.ChainID,
		&i.HaveResults,
		&i.FinalResults,
		&i.ResultsVotes,
		&i.ResultsWeight,
		&i.ResultsBlockHeight,
		&i.CensusRoot,
		&i.MaxCensusSize,
		&i.CensusUri,
		&i.Metadata,
		&i.CensusOrigin,
		&i.Status,
		&i.Namespace,
		&i.Envelope,
		&i.Mode,
		&i.VoteOpts,
		&i.PrivateKeys,
		&i.PublicKeys,
		&i.QuestionIndex,
		&i.CreationTime,
		&i.SourceBlockHeight,
		&i.SourceNetworkID,
		&i.FromArchive,
	)
	return i, err
}

const getProcessCount = `-- name: GetProcessCount :one
SELECT COUNT(*) FROM processes
`

func (q *Queries) GetProcessCount(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getProcessCountStmt, getProcessCount)
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
	rows, err := q.query(ctx, q.getProcessIDsByFinalResultsStmt, getProcessIDsByFinalResults, finalResults)
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
	row := q.queryRow(ctx, q.getProcessStatusStmt, getProcessStatus, id)
	var status int64
	err := row.Scan(&status)
	return status, err
}

const searchEntities = `-- name: SearchEntities :many
SELECT entity_id, COUNT(id) AS process_count FROM processes
WHERE (?1 = '' OR (INSTR(LOWER(HEX(entity_id)), ?1) > 0))
GROUP BY entity_id
ORDER BY creation_time DESC, id ASC
LIMIT ?3
OFFSET ?2
`

type SearchEntitiesParams struct {
	EntityIDSubstr interface{}
	Offset         int64
	Limit          int64
}

type SearchEntitiesRow struct {
	EntityID     types.EntityID
	ProcessCount int64
}

func (q *Queries) SearchEntities(ctx context.Context, arg SearchEntitiesParams) ([]SearchEntitiesRow, error) {
	rows, err := q.query(ctx, q.searchEntitiesStmt, searchEntities, arg.EntityIDSubstr, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchEntitiesRow
	for rows.Next() {
		var i SearchEntitiesRow
		if err := rows.Scan(&i.EntityID, &i.ProcessCount); err != nil {
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

const searchProcesses = `-- name: SearchProcesses :many
SELECT id FROM processes
WHERE (LENGTH(?1) = 0 OR entity_id = ?1)
	AND (?2 = 0 OR namespace = ?2)
	AND (?3 = 0 OR status = ?3)
	AND (?4 = 0 OR source_network_id = ?4)
	-- TODO(mvdan): consider keeping an id_hex column for faster searches
	AND (?5 = '' OR (INSTR(LOWER(HEX(id)), ?5) > 0))
	AND (?6 = FALSE OR have_results)
ORDER BY creation_time DESC, id ASC
LIMIT ?8
OFFSET ?7
`

type SearchProcessesParams struct {
	EntityID        interface{}
	Namespace       interface{}
	Status          interface{}
	SourceNetworkID interface{}
	IDSubstr        interface{}
	WithResults     interface{}
	Offset          int64
	Limit           int64
}

func (q *Queries) SearchProcesses(ctx context.Context, arg SearchProcessesParams) ([]types.ProcessID, error) {
	rows, err := q.query(ctx, q.searchProcessesStmt, searchProcesses,
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
SET have_results = FALSE, final_results = TRUE, 
    end_date = ?1
WHERE id = ?2
`

type SetProcessResultsCancelledParams struct {
	EndDate time.Time
	ID      types.ProcessID
}

func (q *Queries) SetProcessResultsCancelled(ctx context.Context, arg SetProcessResultsCancelledParams) (sql.Result, error) {
	return q.exec(ctx, q.setProcessResultsCancelledStmt, setProcessResultsCancelled, arg.EndDate, arg.ID)
}

const setProcessResultsReady = `-- name: SetProcessResultsReady :execresult
UPDATE processes
SET have_results = TRUE, final_results = TRUE,
	results_votes = ?1,
	results_weight = ?2,
	results_block_height = ?3,
	end_date = ?4
WHERE id = ?5
`

type SetProcessResultsReadyParams struct {
	Votes       string
	Weight      string
	BlockHeight int64
	EndDate     time.Time
	ID          types.ProcessID
}

func (q *Queries) SetProcessResultsReady(ctx context.Context, arg SetProcessResultsReadyParams) (sql.Result, error) {
	return q.exec(ctx, q.setProcessResultsReadyStmt, setProcessResultsReady,
		arg.Votes,
		arg.Weight,
		arg.BlockHeight,
		arg.EndDate,
		arg.ID,
	)
}

const updateProcessEndBlock = `-- name: UpdateProcessEndBlock :execresult
UPDATE processes
SET end_block  = ?1,
	end_date = ?2
WHERE id = ?3
`

type UpdateProcessEndBlockParams struct {
	EndBlock int64
	EndDate  time.Time
	ID       types.ProcessID
}

func (q *Queries) UpdateProcessEndBlock(ctx context.Context, arg UpdateProcessEndBlockParams) (sql.Result, error) {
	return q.exec(ctx, q.updateProcessEndBlockStmt, updateProcessEndBlock, arg.EndBlock, arg.EndDate, arg.ID)
}

const updateProcessFromState = `-- name: UpdateProcessFromState :execresult
;

UPDATE processes
SET census_root         = ?1,
	census_uri          = ?2,
	private_keys        = ?3,
	public_keys         = ?4,
	metadata            = ?5,
	status              = ?6,
	start_date 	        = ?7
WHERE id = ?8
`

type UpdateProcessFromStateParams struct {
	CensusRoot  types.CensusRoot
	CensusUri   string
	PrivateKeys string
	PublicKeys  string
	Metadata    string
	Status      int64
	StartDate   time.Time
	ID          types.ProcessID
}

func (q *Queries) UpdateProcessFromState(ctx context.Context, arg UpdateProcessFromStateParams) (sql.Result, error) {
	return q.exec(ctx, q.updateProcessFromStateStmt, updateProcessFromState,
		arg.CensusRoot,
		arg.CensusUri,
		arg.PrivateKeys,
		arg.PublicKeys,
		arg.Metadata,
		arg.Status,
		arg.StartDate,
		arg.ID,
	)
}

const updateProcessResultByID = `-- name: UpdateProcessResultByID :execresult
UPDATE processes
SET results_votes  = ?1,
    results_weight = ?2,
    vote_opts = ?3,
    envelope = ?4
WHERE id = ?5
`

type UpdateProcessResultByIDParams struct {
	Votes    string
	Weight   string
	VoteOpts string
	Envelope string
	ID       types.ProcessID
}

func (q *Queries) UpdateProcessResultByID(ctx context.Context, arg UpdateProcessResultByIDParams) (sql.Result, error) {
	return q.exec(ctx, q.updateProcessResultByIDStmt, updateProcessResultByID,
		arg.Votes,
		arg.Weight,
		arg.VoteOpts,
		arg.Envelope,
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
	return q.exec(ctx, q.updateProcessResultsStmt, updateProcessResults,
		arg.Votes,
		arg.Weight,
		arg.BlockHeight,
		arg.ID,
	)
}
