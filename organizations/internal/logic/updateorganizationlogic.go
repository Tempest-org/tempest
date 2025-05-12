package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrganizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrganizationLogic {
	return &UpdateOrganizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrganizationLogic) UpdateOrganization(in *organizations.UpdateOrganizationRequest) (*organizations.UpdateOrganizationResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.UpdateOrganizationResponse{}, nil
}
