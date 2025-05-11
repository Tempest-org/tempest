package logic

import (
	"context"

	"github.com/tempest-org/tempest/access/access"
	"github.com/tempest-org/tempest/access/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type HealthCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HealthCheck performs a health check for the service
func (l *HealthCheckLogic) HealthCheck(in *access.HealthCheckRequest) (*access.HealthCheckResponse, error) {
	// Here we could add more sophisticated health checks
	// For now, we just return that the service is serving
	return &access.HealthCheckResponse{
		Status: access.HealthCheckResponse_SERVING,
	}, nil
}