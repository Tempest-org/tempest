package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInviteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteMemberLogic {
	return &InviteMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Member operations
func (l *InviteMemberLogic) InviteMember(in *organizations.InviteMemberRequest) (*organizations.InviteMemberResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.InviteMemberResponse{}, nil
}
