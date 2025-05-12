package logic

import (
	"context"

	"github.com/tempest-org/tempest/organizations/internal/svc"
	"github.com/tempest-org/tempest/organizations/organizations"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveMemberLogic {
	return &RemoveMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveMemberLogic) RemoveMember(in *organizations.RemoveMemberRequest) (*organizations.RemoveMemberResponse, error) {
	// todo: add your logic here and delete this line

	return &organizations.RemoveMemberResponse{}, nil
}
