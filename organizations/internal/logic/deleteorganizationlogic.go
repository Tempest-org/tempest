package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrganizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrganizationLogic {
	return &DeleteOrganizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOrganizationLogic) DeleteOrganization(in *organizations.DeleteOrganizationRequest) (*organizations.DeleteOrganizationResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.DeleteOrganizationResponse{}, nil
}
