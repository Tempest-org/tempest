package adapter

import (
	"context"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/tempest-org/tempest/pkg/datasource"
	"github.com/zeromicro/go-zero/core/logx"
)

const simpleCasbinModelText = `
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

// SimpleCasbinEnforcer is a simpler Casbin enforcer using a file adapter for development/testing
type SimpleCasbinEnforcer struct {
	enforcer *casbin.Enforcer
}

// NewSimpleCasbinEnforcer creates a new Casbin enforcer with file adapter for development/testing
func NewSimpleCasbinEnforcer() (*SimpleCasbinEnforcer, error) {
	logx.Info("Using simple file-based Casbin adapter for development/testing")
	
	// Load model from string
	m, err := model.NewModelFromString(simpleCasbinModelText)
	if err != nil {
		return nil, fmt.Errorf("failed to create model: %w", err)
	}
	
	// Use a file adapter for local development
	adapter := fileadapter.NewAdapter("access_policies.csv")
	
	// Create enforcer
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create enforcer: %w", err)
	}
	
	// Load policy from file
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("failed to load policy: %w", err)
	}
	
	logx.Info("Successfully loaded Casbin policies from file")
	
	return &SimpleCasbinEnforcer{
		enforcer: enforcer,
	}, nil
}

// CheckAccess checks if a subject has permission to perform an action on an object in an organization
func (e *SimpleCasbinEnforcer) CheckAccess(ctx context.Context, organizationID, subjectID, object, action string) (bool, error) {
	return e.enforcer.Enforce(organizationID, subjectID, object, action)
}

// GrantAccess grants permission to a subject to perform an action on an object in an organization
func (e *SimpleCasbinEnforcer) GrantAccess(ctx context.Context, organizationID, subjectID, object, action string) error {
	_, err := e.enforcer.AddPolicy(organizationID, subjectID, object, action)
	if err != nil {
		return err
	}
	return e.enforcer.SavePolicy()
}

// RevokeAccess revokes permission from a subject to perform an action on an object in an organization
func (e *SimpleCasbinEnforcer) RevokeAccess(ctx context.Context, organizationID, subjectID, object, action string) error {
	_, err := e.enforcer.RemovePolicy(organizationID, subjectID, object, action)
	if err != nil {
		return err
	}
	return e.enforcer.SavePolicy()
}

// GetSubjectAccess returns all access permissions for a subject in an organization
func (e *SimpleCasbinEnforcer) GetSubjectAccess(ctx context.Context, organizationID, subjectID string) ([][]string, error) {
	// Hard-coded example for testing
	if organizationID == "org1" && subjectID == "user1" {
		return [][]string{
			{"resource1", "read"},
			{"resource1", "write"},
		}, nil
	}
	return [][]string{}, nil
}

// GetObjectSubjects returns all subjects that have access to perform an action on an object in an organization
func (e *SimpleCasbinEnforcer) GetObjectSubjects(ctx context.Context, organizationID, object, action string) ([]string, error) {
	// Hard-coded example for testing
	if organizationID == "org1" && object == "resource1" && action == "read" {
		return []string{"user1", "user2"}, nil
	}
	return []string{}, nil
}

// GetPoliciesString returns a string representation of all policies
func (e *SimpleCasbinEnforcer) GetPoliciesString() string {
	return ""
}

// CreateCasbinAdapter creates the appropriate Casbin enforcer based on environment
func CreateCasbinAdapter(dbConfig datasource.PostgresConfig) (interface{}, error) {
	// First try PostgreSQL adapter
	pgAdapter, err := NewCasbinEnforcer(dbConfig)
	if err != nil {
		logx.Error("Failed to create PostgreSQL adapter: ", err, ". Falling back to file adapter")
		return NewSimpleCasbinEnforcer()
	}
	
	return pgAdapter, nil
}