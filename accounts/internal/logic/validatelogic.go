package logic

import (
	"context"

	"github.com/tempest-org/tempest/accounts/accounts"
	"github.com/tempest-org/tempest/accounts/internal/model"
	"github.com/tempest-org/tempest/accounts/internal/svc"
	"github.com/tempest-org/tempest/pkg/datasource"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	tokens TokensLogic
	model model.AccountsModel
}

func NewValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateLogic {
	logger := logx.WithContext(ctx)
	tokens, err := NewTokensLogic()
	if err != nil {
		logger.Error(err)
		return nil
	}
	pgConn := datasource.NewPostgresConn(svcCtx.Config.Database.URI)
	accModel := model.NewAccountsModel(pgConn)

	return &ValidateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		tokens: tokens,
		model:  accModel,
	}
}

func (l *ValidateLogic) Validate(in *accounts.ValidateRequest) (*accounts.ValidateResponse, error) {
	res := &accounts.ValidateResponse{
		Valid: false,
	}

	t, err := l.tokens.Parse(l.ctx, in.AccessToken)
	if err != nil {
		l.Logger.Error("Could not parse access token", "err", err.Error())
		return res, err
	}

		var username string
		err= t.Get("username", &username)
		if err != nil {
			l.Logger.Error("Could not get username from access token", "err", err.Error())
			return res, err
		}

		var email string
		err =t.Get("email", &email)
		if err != nil {
			l.Logger.Error("Could not get email from access token", "err", err.Error())
			return res, err
		}

		claims := map[string]string{
			"username": username,
			"email": email,
		}

		sub, _ := t.Subject()
		e, _ := t.Expiration()
		exp := e.Unix()
		issuer, _ := t.Issuer()

	return &accounts.ValidateResponse{
		Valid: true,
		Sub: &sub,
		Exp: &exp,
		Issuer: &issuer,
		Claims: claims,
	}, nil
}
