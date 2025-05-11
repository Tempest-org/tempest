package logic

import (
	"context"

	"github.com/tempest-org/tempest/accounts/accounts"
	"github.com/tempest-org/tempest/accounts/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *accounts.Request) (*accounts.Response, error) {
	// todo: add your logic here and delete this line

	return &accounts.Response{}, nil
}
