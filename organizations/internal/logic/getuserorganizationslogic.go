package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserOrganizationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserOrganizationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOrganizationsLogic {
	return &GetUserOrganizationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// User-specific operations
func (l *GetUserOrganizationsLogic) GetUserOrganizations(in *organizations.GetUserOrganizationsRequest) (*organizations.GetUserOrganizationsResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.GetUserOrganizationsResponse{}, nil
}
