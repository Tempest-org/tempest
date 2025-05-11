package logic

import (
	"context"

	"accounts/accounts"
	"accounts/internal/model"
	"accounts/internal/svc"
	"accounts/pkg/datasource"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext

	tokens TokensLogic
	model model.AccountsModel

	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	logger:= logx.WithContext(ctx)
	tokens , err := NewTokensLogic()
	if err != nil {
		logger.Error(err)
		return nil
	}
	pgConn := datasource.NewPostgresConn(svcCtx.Config.Database.URI)
	accModel := model.NewAccountsModel(pgConn)
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger,
		tokens: tokens,
		model:  accModel,
	}
}

func (l *LoginLogic) Login(in *accounts.LoginRequest) (*accounts.TokenResponse, error) {
	// todo: add your logic here and delete this line

	acc, err := l.accountByIdentifier(in.Identifier, in.Value)
	if err != nil {
		l.Logger.Error("Could not find account", "err", err.Error())
		return nil, err
	}

	if acc == nil {
		l.Logger.Error("Could not find account")
		return nil, err
	}

	err= bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(in.Password))
	if err != nil {
		l.Logger.Error("Invalid password", "err", err.Error())
		return nil, err
	}

	extraClaims := map[string]interface{}{
		"username": acc.Username,
		"email":    acc.Email,
	}

	access, refresh, err := l.tokens.GenTokenPair(l.ctx, acc.Id, extraClaims)
	if err != nil {
		l.Logger.Error("Could not generate tokens", "err", err.Error())
		return nil, err
	}

	return &accounts.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (l *LoginLogic) accountByIdentifier (identifier accounts.Identifier, value string) (*model.Accounts, error) {
	switch (identifier){
		case *accounts.Identifier_EMAIL.Enum():
			return l.model.FindOneByEmail(context.Background(), value)
		case *accounts.Identifier_USERNAME.Enum():
			return l.model.FindOneByUsername(context.Background(), value)
		case *accounts.Identifier_PHONE.Enum():
			return l.model.FindOneByPhone(context.Background(), value)
		default:
			return nil, nil
	}
}
