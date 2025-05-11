package adapter

import (
	"context"
	"fmt"

	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tempest-org/tempest/pkg/datasource"
	"github.com/zeromicro/go-zero/core/logx"
)

const casbinModelText = `
[request_definition]
r = org, sub, obj, act

[policy_definition]
p = org, sub, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.org == p.org && r.sub == p.sub && r.obj == p.obj && r.act == p.act
`

type CasbinEnforcer struct {
	enforcer *casbin.Enforcer
}

// Creates a new Casbin enforcer with the specified database configuration
func NewCasbinEnforcer(dbConfig datasource.PostgresConfig) (*CasbinEnforcer, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	
	logx.Infof("Connecting to PostgreSQL at %s:%d/%s", dbConfig.Host, dbConfig.Port, dbConfig.Database)

	adapter, err := pgadapter.NewAdapter(connString)
	if err != nil {
		logx.Errorf("Failed to create adapter: %v", err)
		return nil, fmt.Errorf("failed to create adapter: %w", err)
	}
	
	logx.Info("Successfully created Casbin PostgreSQL adapter")

	m, err := model.NewModelFromString(casbinModelText)
	if err != nil {
		return nil, fmt.Errorf("failed to create model: %w", err)
	}

	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create enforcer: %w", err)
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("failed to load policy: %w", err)
	}
	
	logx.Info("Successfully loaded Casbin policies")

	return &CasbinEnforcer{
		enforcer: enforcer,
	}, nil
}

// CheckAccess checks if a subject has permission to perform an action on an object in an organization
func (e *CasbinEnforcer) CheckAccess(ctx context.Context, organizationID, subjectID, object, action string) (bool, error) {
	return e.enforcer.Enforce(organizationID, subjectID, object, action)
}

// GrantAccess grants permission to a subject to perform an action on an object in an organization
func (e *CasbinEnforcer) GrantAccess(ctx context.Context, organizationID, subjectID, object, action string) error {
	_, err := e.enforcer.AddPolicy(organizationID, subjectID, object, action)
	return err
}

// RevokeAccess revokes permission from a subject to perform an action on an object in an organization
func (e *CasbinEnforcer) RevokeAccess(ctx context.Context, organizationID, subjectID, object, action string) error {
	_, err := e.enforcer.RemovePolicy(organizationID, subjectID, object, action)
	return err
}

// GetSubjectAccess returns all access permissions for a subject in an organization
func (e *CasbinEnforcer) GetSubjectAccess(ctx context.Context, organizationID, subjectID string) ([][]string, error) {
	return e.enforcer.GetFilteredPolicy(0, organizationID, subjectID)
}

// GetObjectSubjects returns all subjects that have access to perform an action on an object in an organization
func (e *CasbinEnforcer) GetObjectSubjects(ctx context.Context, organizationID, object, action string) ([]string, error) {
	policies,err := e.enforcer.GetFilteredPolicy(0, organizationID)
	if err != nil {
		return nil, err
	}
	subjects := make([]string, 0)

	for _, policy := range policies {
		if len(policy) >= 4 && policy[2] == object && policy[3] == action {
			subjects = append(subjects, policy[1])
		}
	}

	return subjects, nil
}

// EnsureDBTables ensures that all necessary tables are created in the database
func EnsureDBTables(pool *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS casbin_rule (
		id SERIAL PRIMARY KEY,
		ptype VARCHAR(100) NOT NULL,
		v0 VARCHAR(100),
		v1 VARCHAR(100),
		v2 VARCHAR(100),
		v3 VARCHAR(100),
		v4 VARCHAR(100),
		v5 VARCHAR(100)
	);
	`
	_, err := pool.Exec(context.Background(), query)
	return err
}
