package logic

import (
	"context"

	"accounts/accounts"
	"accounts/internal/model"
	"accounts/internal/svc"
	"accounts/pkg/datasource"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	tokens TokensLogic
	model model.AccountsModel
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	logger:= logx.WithContext(ctx)
	tokens, err := NewTokensLogic()
	if err != nil {
		logger.Error(err)
		return nil
	}
	pgConn := datasource.NewPostgresConn(svcCtx.Config.Database.URI)
	accModel := model.NewAccountsModel(pgConn)
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		tokens: tokens,
		model:  accModel,
	}
}

func (l *RefreshLogic) Refresh(in *accounts.RefreshRequest) (*accounts.RefreshResponse, error) {
	token, err := l.tokens.Parse(l.ctx, in.RefreshToken)
	if err != nil {
		l.Logger.Error("Could not parse refresh token", "err", err.Error())
		return nil, err
	}

	var email string
	err = token.Get("email", &email)
	if err != nil {
		l.Logger.Error("Could not get email from refresh token", "err", err.Error())
		return nil, err
	}

	acc, err := l.model.FindOneByEmail(l.ctx, email)
	if err != nil {
		l.Logger.Error("Could not find account", "err", err.Error())
		return nil, err
	}

	extraClaims := map[string]interface{}{
		"email": email,
		"username": acc.Username,
	}

	p := l.tokens.NewTokenPayload(acc.Id, "access", extraClaims)
	at, err := l.tokens.Build(l.ctx, p)
	if err != nil {
		l.Logger.Error("Could not build refresh token", "err", err.Error())
		return nil, err
	}
	access, err := l.tokens.Sign(l.ctx, at)
	if err != nil {
		l.Logger.Error("Could not sign refresh token", "err", err.Error())
		return nil, err
	}

	return &accounts.RefreshResponse{
		AccessToken: access,
	}, nil
}
