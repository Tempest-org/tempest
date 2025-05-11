package svc

import (
	"context"

	"github.com/tempest-org/tempest/access/internal/adapter"
	"github.com/tempest-org/tempest/access/internal/config"
	"github.com/tempest-org/tempest/pkg/datasource"
	"github.com/zeromicro/go-zero/core/logx"
)

// Define interface for Casbin enforcement
type AccessEnforcer interface {
	CheckAccess(ctx context.Context, organizationID, subjectID, object, action string) (bool, error)
	GrantAccess(ctx context.Context, organizationID, subjectID, object, action string) error
	RevokeAccess(ctx context.Context, organizationID, subjectID, object, action string) error
	GetSubjectAccess(ctx context.Context, organizationID, subjectID string) ([][]string, error)
	GetObjectSubjects(ctx context.Context, organizationID, object, action string) ([]string, error)
}

type ServiceContext struct {
	Config    config.Config
	Enforcer  AccessEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Create database configuration
	dbConfig := datasource.PostgresConfig{
		Host:     c.Database.Host,
		Port:     c.Database.Port,
		User:     c.Database.User,
		Password: c.Database.Password,
		Database: c.Database.Database,
	}

	// Try to create Casbin adapter, fallback to simple adapter if needed
	enforcer, err := adapter.CreateCasbinAdapter(dbConfig)
	if err != nil {
		logx.Errorf("Failed to create access enforcer: %v", err)
		return nil
	}

	// Check which type of enforcer we got
	var accessEnforcer AccessEnforcer
	if pgEnforcer, ok := enforcer.(*adapter.CasbinEnforcer); ok {
		accessEnforcer = pgEnforcer
		logx.Info("Using PostgreSQL Casbin adapter")
	} else if simpleEnforcer, ok := enforcer.(*adapter.SimpleCasbinEnforcer); ok {
		accessEnforcer = simpleEnforcer
		logx.Info("Using file-based Casbin adapter")
	} else {
		logx.Error("Unknown enforcer type")
		return nil
	}

	return &ServiceContext{
		Config:    c,
		Enforcer:  accessEnforcer,
	}
}
