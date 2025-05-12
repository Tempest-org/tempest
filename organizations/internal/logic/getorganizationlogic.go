package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrganizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrganizationLogic {
	return &GetOrganizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrganizationLogic) GetOrganization(in *organizations.GetOrganizationRequest) (*organizations.GetOrganizationResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.GetOrganizationResponse{}, nil
}
