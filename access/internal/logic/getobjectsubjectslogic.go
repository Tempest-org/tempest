package logic

import (
	"context"

	"github.com/tempest-org/tempest/access/access"
	"github.com/tempest-org/tempest/access/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetObjectSubjectsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetObjectSubjectsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetObjectSubjectsLogic {
	return &GetObjectSubjectsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetObjectSubjects returns all subjects that have access to perform an action on an object in an organization
func (l *GetObjectSubjectsLogic) GetObjectSubjects(in *access.GetObjectSubjectsRequest) (*access.GetObjectSubjectsResponse, error) {
	subjects, err := l.svcCtx.Enforcer.GetObjectSubjects(
		l.ctx,
		in.OrganizationId,
		in.Object,
		in.Action,
	)
	if err != nil {
		return nil, err
	}

	return &access.GetObjectSubjectsResponse{
		SubjectIds: subjects,
	}, nil
}