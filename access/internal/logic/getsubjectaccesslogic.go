package logic

import (
	"context"

	"github.com/tempest-org/tempest/access/access"
	"github.com/tempest-org/tempest/access/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubjectAccessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubjectAccessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubjectAccessLogic {
	return &GetSubjectAccessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetSubjectAccess returns all access permissions for a subject in an organization
func (l *GetSubjectAccessLogic) GetSubjectAccess(in *access.GetSubjectAccessRequest) (*access.GetSubjectAccessResponse, error) {
	policies, err := l.svcCtx.Enforcer.GetSubjectAccess(
		l.ctx,
		in.OrganizationId,
		in.SubjectId,
	)
	if err != nil {
		return nil, err
	}

	permissions := make([]*access.AccessPermission, 0, len(policies))
	for _, policy := range policies {
		if len(policy) >= 4 {
			// Format: [org, subject, object, action]
			permissions = append(permissions, &access.AccessPermission{
				Object: policy[2],
				Action: policy[3],
			})
		}
	}

	return &access.GetSubjectAccessResponse{
		Permissions: permissions,
	}, nil
}