package datasource

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const postgresDriverName = "pgx"

func NewPostgresConn(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(postgresDriverName, datasource, opts...)
}
