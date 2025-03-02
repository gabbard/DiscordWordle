// Code generated by sqlc. DO NOT EDIT.

package wordle

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.checkIfServerHasDisabledQuipsStmt, err = db.PrepareContext(ctx, checkIfServerHasDisabledQuips); err != nil {
		return nil, fmt.Errorf("error preparing query CheckIfServerHasDisabledQuips: %w", err)
	}
	if q.countAccountsByDiscordIdStmt, err = db.PrepareContext(ctx, countAccountsByDiscordId); err != nil {
		return nil, fmt.Errorf("error preparing query CountAccountsByDiscordId: %w", err)
	}
	if q.countNicknameByDiscordIdAndServerIdStmt, err = db.PrepareContext(ctx, countNicknameByDiscordIdAndServerId); err != nil {
		return nil, fmt.Errorf("error preparing query CountNicknameByDiscordIdAndServerId: %w", err)
	}
	if q.countScoresByDiscordIdStmt, err = db.PrepareContext(ctx, countScoresByDiscordId); err != nil {
		return nil, fmt.Errorf("error preparing query CountScoresByDiscordId: %w", err)
	}
	if q.createAccountStmt, err = db.PrepareContext(ctx, createAccount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccount: %w", err)
	}
	if q.createNicknameStmt, err = db.PrepareContext(ctx, createNickname); err != nil {
		return nil, fmt.Errorf("error preparing query CreateNickname: %w", err)
	}
	if q.createQuipForScoreStmt, err = db.PrepareContext(ctx, createQuipForScore); err != nil {
		return nil, fmt.Errorf("error preparing query CreateQuipForScore: %w", err)
	}
	if q.createScoreStmt, err = db.PrepareContext(ctx, createScore); err != nil {
		return nil, fmt.Errorf("error preparing query CreateScore: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.deleteNicknameStmt, err = db.PrepareContext(ctx, deleteNickname); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteNickname: %w", err)
	}
	if q.deleteScoresForUserStmt, err = db.PrepareContext(ctx, deleteScoresForUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteScoresForUser: %w", err)
	}
	if q.disableQuipsForServerStmt, err = db.PrepareContext(ctx, disableQuipsForServer); err != nil {
		return nil, fmt.Errorf("error preparing query DisableQuipsForServer: %w", err)
	}
	if q.enableQuipsForServerStmt, err = db.PrepareContext(ctx, enableQuipsForServer); err != nil {
		return nil, fmt.Errorf("error preparing query EnableQuipsForServer: %w", err)
	}
	if q.getAccountStmt, err = db.PrepareContext(ctx, getAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccount: %w", err)
	}
	if q.getExpectedPreviousWeekGamesStmt, err = db.PrepareContext(ctx, getExpectedPreviousWeekGames); err != nil {
		return nil, fmt.Errorf("error preparing query GetExpectedPreviousWeekGames: %w", err)
	}
	if q.getExpectedWeekGamesStmt, err = db.PrepareContext(ctx, getExpectedWeekGames); err != nil {
		return nil, fmt.Errorf("error preparing query GetExpectedWeekGames: %w", err)
	}
	if q.getNicknameStmt, err = db.PrepareContext(ctx, getNickname); err != nil {
		return nil, fmt.Errorf("error preparing query GetNickname: %w", err)
	}
	if q.getNicknamesByDiscordIdStmt, err = db.PrepareContext(ctx, getNicknamesByDiscordId); err != nil {
		return nil, fmt.Errorf("error preparing query GetNicknamesByDiscordId: %w", err)
	}
	if q.getQuipByScoreStmt, err = db.PrepareContext(ctx, getQuipByScore); err != nil {
		return nil, fmt.Errorf("error preparing query GetQuipByScore: %w", err)
	}
	if q.getQuipsByCreatedByAccountStmt, err = db.PrepareContext(ctx, getQuipsByCreatedByAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetQuipsByCreatedByAccount: %w", err)
	}
	if q.getScoreHistoryByAccountStmt, err = db.PrepareContext(ctx, getScoreHistoryByAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetScoreHistoryByAccount: %w", err)
	}
	if q.getScoresByServerIdStmt, err = db.PrepareContext(ctx, getScoresByServerId); err != nil {
		return nil, fmt.Errorf("error preparing query GetScoresByServerId: %w", err)
	}
	if q.getScoresByServerIdPreviousWeekStmt, err = db.PrepareContext(ctx, getScoresByServerIdPreviousWeek); err != nil {
		return nil, fmt.Errorf("error preparing query GetScoresByServerIdPreviousWeek: %w", err)
	}
	if q.incrementQuipStmt, err = db.PrepareContext(ctx, incrementQuip); err != nil {
		return nil, fmt.Errorf("error preparing query IncrementQuip: %w", err)
	}
	if q.listAccountsStmt, err = db.PrepareContext(ctx, listAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query ListAccounts: %w", err)
	}
	if q.listNicknamesStmt, err = db.PrepareContext(ctx, listNicknames); err != nil {
		return nil, fmt.Errorf("error preparing query ListNicknames: %w", err)
	}
	if q.listScoresStmt, err = db.PrepareContext(ctx, listScores); err != nil {
		return nil, fmt.Errorf("error preparing query ListScores: %w", err)
	}
	if q.updateNicknameStmt, err = db.PrepareContext(ctx, updateNickname); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateNickname: %w", err)
	}
	if q.updateScoreStmt, err = db.PrepareContext(ctx, updateScore); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateScore: %w", err)
	}
	if q.updateTimeZoneStmt, err = db.PrepareContext(ctx, updateTimeZone); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTimeZone: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.checkIfServerHasDisabledQuipsStmt != nil {
		if cerr := q.checkIfServerHasDisabledQuipsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing checkIfServerHasDisabledQuipsStmt: %w", cerr)
		}
	}
	if q.countAccountsByDiscordIdStmt != nil {
		if cerr := q.countAccountsByDiscordIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countAccountsByDiscordIdStmt: %w", cerr)
		}
	}
	if q.countNicknameByDiscordIdAndServerIdStmt != nil {
		if cerr := q.countNicknameByDiscordIdAndServerIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countNicknameByDiscordIdAndServerIdStmt: %w", cerr)
		}
	}
	if q.countScoresByDiscordIdStmt != nil {
		if cerr := q.countScoresByDiscordIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countScoresByDiscordIdStmt: %w", cerr)
		}
	}
	if q.createAccountStmt != nil {
		if cerr := q.createAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountStmt: %w", cerr)
		}
	}
	if q.createNicknameStmt != nil {
		if cerr := q.createNicknameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createNicknameStmt: %w", cerr)
		}
	}
	if q.createQuipForScoreStmt != nil {
		if cerr := q.createQuipForScoreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createQuipForScoreStmt: %w", cerr)
		}
	}
	if q.createScoreStmt != nil {
		if cerr := q.createScoreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createScoreStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.deleteNicknameStmt != nil {
		if cerr := q.deleteNicknameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteNicknameStmt: %w", cerr)
		}
	}
	if q.deleteScoresForUserStmt != nil {
		if cerr := q.deleteScoresForUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteScoresForUserStmt: %w", cerr)
		}
	}
	if q.disableQuipsForServerStmt != nil {
		if cerr := q.disableQuipsForServerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing disableQuipsForServerStmt: %w", cerr)
		}
	}
	if q.enableQuipsForServerStmt != nil {
		if cerr := q.enableQuipsForServerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing enableQuipsForServerStmt: %w", cerr)
		}
	}
	if q.getAccountStmt != nil {
		if cerr := q.getAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountStmt: %w", cerr)
		}
	}
	if q.getExpectedPreviousWeekGamesStmt != nil {
		if cerr := q.getExpectedPreviousWeekGamesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getExpectedPreviousWeekGamesStmt: %w", cerr)
		}
	}
	if q.getExpectedWeekGamesStmt != nil {
		if cerr := q.getExpectedWeekGamesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getExpectedWeekGamesStmt: %w", cerr)
		}
	}
	if q.getNicknameStmt != nil {
		if cerr := q.getNicknameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNicknameStmt: %w", cerr)
		}
	}
	if q.getNicknamesByDiscordIdStmt != nil {
		if cerr := q.getNicknamesByDiscordIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getNicknamesByDiscordIdStmt: %w", cerr)
		}
	}
	if q.getQuipByScoreStmt != nil {
		if cerr := q.getQuipByScoreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getQuipByScoreStmt: %w", cerr)
		}
	}
	if q.getQuipsByCreatedByAccountStmt != nil {
		if cerr := q.getQuipsByCreatedByAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getQuipsByCreatedByAccountStmt: %w", cerr)
		}
	}
	if q.getScoreHistoryByAccountStmt != nil {
		if cerr := q.getScoreHistoryByAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScoreHistoryByAccountStmt: %w", cerr)
		}
	}
	if q.getScoresByServerIdStmt != nil {
		if cerr := q.getScoresByServerIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScoresByServerIdStmt: %w", cerr)
		}
	}
	if q.getScoresByServerIdPreviousWeekStmt != nil {
		if cerr := q.getScoresByServerIdPreviousWeekStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getScoresByServerIdPreviousWeekStmt: %w", cerr)
		}
	}
	if q.incrementQuipStmt != nil {
		if cerr := q.incrementQuipStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing incrementQuipStmt: %w", cerr)
		}
	}
	if q.listAccountsStmt != nil {
		if cerr := q.listAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAccountsStmt: %w", cerr)
		}
	}
	if q.listNicknamesStmt != nil {
		if cerr := q.listNicknamesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listNicknamesStmt: %w", cerr)
		}
	}
	if q.listScoresStmt != nil {
		if cerr := q.listScoresStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listScoresStmt: %w", cerr)
		}
	}
	if q.updateNicknameStmt != nil {
		if cerr := q.updateNicknameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateNicknameStmt: %w", cerr)
		}
	}
	if q.updateScoreStmt != nil {
		if cerr := q.updateScoreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateScoreStmt: %w", cerr)
		}
	}
	if q.updateTimeZoneStmt != nil {
		if cerr := q.updateTimeZoneStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTimeZoneStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                      DBTX
	tx                                      *sql.Tx
	checkIfServerHasDisabledQuipsStmt       *sql.Stmt
	countAccountsByDiscordIdStmt            *sql.Stmt
	countNicknameByDiscordIdAndServerIdStmt *sql.Stmt
	countScoresByDiscordIdStmt              *sql.Stmt
	createAccountStmt                       *sql.Stmt
	createNicknameStmt                      *sql.Stmt
	createQuipForScoreStmt                  *sql.Stmt
	createScoreStmt                         *sql.Stmt
	deleteAccountStmt                       *sql.Stmt
	deleteNicknameStmt                      *sql.Stmt
	deleteScoresForUserStmt                 *sql.Stmt
	disableQuipsForServerStmt               *sql.Stmt
	enableQuipsForServerStmt                *sql.Stmt
	getAccountStmt                          *sql.Stmt
	getExpectedPreviousWeekGamesStmt        *sql.Stmt
	getExpectedWeekGamesStmt                *sql.Stmt
	getNicknameStmt                         *sql.Stmt
	getNicknamesByDiscordIdStmt             *sql.Stmt
	getQuipByScoreStmt                      *sql.Stmt
	getQuipsByCreatedByAccountStmt          *sql.Stmt
	getScoreHistoryByAccountStmt            *sql.Stmt
	getScoresByServerIdStmt                 *sql.Stmt
	getScoresByServerIdPreviousWeekStmt     *sql.Stmt
	incrementQuipStmt                       *sql.Stmt
	listAccountsStmt                        *sql.Stmt
	listNicknamesStmt                       *sql.Stmt
	listScoresStmt                          *sql.Stmt
	updateNicknameStmt                      *sql.Stmt
	updateScoreStmt                         *sql.Stmt
	updateTimeZoneStmt                      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                      tx,
		tx:                                      tx,
		checkIfServerHasDisabledQuipsStmt:       q.checkIfServerHasDisabledQuipsStmt,
		countAccountsByDiscordIdStmt:            q.countAccountsByDiscordIdStmt,
		countNicknameByDiscordIdAndServerIdStmt: q.countNicknameByDiscordIdAndServerIdStmt,
		countScoresByDiscordIdStmt:              q.countScoresByDiscordIdStmt,
		createAccountStmt:                       q.createAccountStmt,
		createNicknameStmt:                      q.createNicknameStmt,
		createQuipForScoreStmt:                  q.createQuipForScoreStmt,
		createScoreStmt:                         q.createScoreStmt,
		deleteAccountStmt:                       q.deleteAccountStmt,
		deleteNicknameStmt:                      q.deleteNicknameStmt,
		deleteScoresForUserStmt:                 q.deleteScoresForUserStmt,
		disableQuipsForServerStmt:               q.disableQuipsForServerStmt,
		enableQuipsForServerStmt:                q.enableQuipsForServerStmt,
		getAccountStmt:                          q.getAccountStmt,
		getExpectedPreviousWeekGamesStmt:        q.getExpectedPreviousWeekGamesStmt,
		getExpectedWeekGamesStmt:                q.getExpectedWeekGamesStmt,
		getNicknameStmt:                         q.getNicknameStmt,
		getNicknamesByDiscordIdStmt:             q.getNicknamesByDiscordIdStmt,
		getQuipByScoreStmt:                      q.getQuipByScoreStmt,
		getQuipsByCreatedByAccountStmt:          q.getQuipsByCreatedByAccountStmt,
		getScoreHistoryByAccountStmt:            q.getScoreHistoryByAccountStmt,
		getScoresByServerIdStmt:                 q.getScoresByServerIdStmt,
		getScoresByServerIdPreviousWeekStmt:     q.getScoresByServerIdPreviousWeekStmt,
		incrementQuipStmt:                       q.incrementQuipStmt,
		listAccountsStmt:                        q.listAccountsStmt,
		listNicknamesStmt:                       q.listNicknamesStmt,
		listScoresStmt:                          q.listScoresStmt,
		updateNicknameStmt:                      q.updateNicknameStmt,
		updateScoreStmt:                         q.updateScoreStmt,
		updateTimeZoneStmt:                      q.updateTimeZoneStmt,
	}
}
