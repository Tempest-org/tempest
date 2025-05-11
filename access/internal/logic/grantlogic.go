package logic

import (
	"context"

	"github.com/tempest-org/tempest/access/access"
	"github.com/tempest-org/tempest/access/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GrantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrantLogic {
	return &GrantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Grant grants permission to a subject to perform an action on an object in an organization
func (l *GrantLogic) Grant(in *access.GrantAccessRequest) (*access.Empty, error) {
	err := l.svcCtx.Enforcer.GrantAccess(
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