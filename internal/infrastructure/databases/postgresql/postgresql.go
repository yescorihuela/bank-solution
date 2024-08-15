package postgresql

import (
	"context"
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/shared/utils"
)

type PGXQueryer interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func WithTX(ctx context.Context, conn *pgxpool.Pool, f func(sTx pgx.Tx) error) error {
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return fmt.Errorf("error on begin %w", err)
	}

	if err := f(tx); err != nil {
		_ = tx.Rollback(ctx)

		return fmt.Errorf("rolling back transaction %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error on commit %w", err)
	}

	return nil
}

func NewPostgresDBConnection(config utils.Config) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.BlueSoftURLDB)
	if err != nil {
		return nil, err
	}

	if config.MaxDBConnections <= 0 {
		config.MaxDBConnections = 10
	}
	cfg.MinConns = 0
	cfg.MaxConnIdleTime = 30 * time.Minute
	cfg.MaxConnLifetime = 60 * time.Minute

	cfg.ConnConfig.LogLevel = pgx.LogLevelDebug

	return pgxpool.ConnectConfig(context.Background(), cfg)
}
