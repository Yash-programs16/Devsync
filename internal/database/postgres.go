package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Yash-programs16/devsync/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Database wraps the PostgreSQL connection pool.
//
// This allows us to add helper methods later such as:
//   - Health()
//   - Stats()
//   - BeginTx()
//   - Repository constructors
type Database struct {
	Pool *pgxpool.Pool
}

// New creates a PostgreSQL connection pool and verifies connectivity.
func New(cfg *configs.Config) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse postgres config: %w", err)
	}

	// Connection Pool Configuration
	poolConfig.MaxConns = cfg.DBMaxConns
	poolConfig.MinConns = cfg.DBMinConns

	// Recommended production defaults
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 15 * time.Minute
	poolConfig.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &Database{
		Pool: pool,
	}, nil
}

// Ping verifies that the database is reachable.
func (db *Database) Ping(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}

// Close gracefully closes all pooled connections.
func (db *Database) Close() {
	if db != nil && db.Pool != nil {
		db.Pool.Close()
	}
}
