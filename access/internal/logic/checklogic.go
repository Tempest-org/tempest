package logic

import (
	"context"

	"github.com/tempest-org/tempest/access/access"
	"github.com/tempest-org/tempest/access/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *access.CheckAccessRequest) (*access.CheckAccessResponse, error) {
	l.Logger.Infof("Checking if subject %s has permission to perform action %s on object %s in organization %s", in.SubjectId, in.Action, in.Object, in.OrganizationId)
	allowed, err := l.svcCtx.Enforcer.CheckAccess(
		l.ctx,
		in.OrganizationId,
		in.SubjectId,
		in.Object,
		in.Action,
	)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	l.Logger.Info("Check result: ", allowed)

	res := &access.CheckAccessResponse{
		Allowed: allowed,
	}

	return res, nil
}
