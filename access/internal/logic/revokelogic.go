package logic

import (
	"context"

	"github.com/tempest-org/tempest/access/access"
	"github.com/tempest-org/tempest/access/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type RevokeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRevokeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevokeLogic {
	return &RevokeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Revoke revokes permission from a subject to perform an action on an object in an organization
func (l *RevokeLogic) Revoke(in *access.RevokeAccessRequest) (*access.Empty, error) {
	err := l.svcCtx.Enforcer.RevokeAccess(
		l.ctx,
		in.OrganizationId,
		in.SubjectId,
		in.Object,
		in.Action,
	)
	if err != nil {
		return nil, err
	}

	return &access.Empty{}, nil
}