package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInvitationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInvitationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInvitationsLogic {
	return &GetInvitationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInvitationsLogic) GetInvitations(in *organizations.GetInvitationsRequest) (*organizations.GetInvitationsResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.GetInvitationsResponse{}, nil
}
