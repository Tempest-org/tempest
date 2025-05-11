// Package datasource provides database connection and query functionality
package datasource

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const postgresDriverName = "pgx"

// PostgresConfig contains database connection parameters
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// Creates a new Postgres connection pool using pgxpool
func NewPostgresConnectionPool(cfg PostgresConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	return pgxpool.New(context.Background(), connStr)
}

// Creates a new Postgres connection using go-zero's sqlx
func NewPostgresConn(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(postgresDriverName, datasource, opts...)
}

// Formats a PostgreSQL connection string from config
func GetConnectionString(cfg PostgresConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
