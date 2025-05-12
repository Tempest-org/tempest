package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrganizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrganizationLogic {
	return &CreateOrganizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Organization operations
func (l *CreateOrganizationLogic) CreateOrganization(in *organizations.CreateOrganizationRequest) (*organizations.CreateOrganizationResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.CreateOrganizationResponse{}, nil
}
