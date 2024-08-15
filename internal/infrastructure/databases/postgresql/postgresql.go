package postgresql

import (
	"context"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yescorihuela/bluesoft-bank-solution/internal/shared/utils"
)

type PGXQueryer interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
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
