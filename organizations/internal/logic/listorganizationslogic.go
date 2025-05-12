package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrganizationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListOrganizationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrganizationsLogic {
	return &ListOrganizationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListOrganizationsLogic) ListOrganizations(in *organizations.ListOrganizationsRequest) (*organizations.ListOrganizationsResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.ListOrganizationsResponse{}, nil
}
