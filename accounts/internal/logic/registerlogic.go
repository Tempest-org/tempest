package logic

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"

	"accounts/accounts"
	"accounts/internal/model"
	"accounts/internal/svc"
	"accounts/pkg/datasource"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	tokens TokensLogic
	model model.AccountsModel

	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	logger:= logx.WithContext(ctx)
	tokens , err := NewTokensLogic()
	pgConn := datasource.NewPostgresConn(svcCtx.Config.Database.URI)
	accModel := model.NewAccountsModel(pgConn)

	if err != nil {
		logger.Error(err)
		return nil
	}

	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger,
		tokens: tokens,
		model:  accModel,
	}
}

func (l *RegisterLogic) Register(in *accounts.RegisterRequest) (*accounts.TokenResponse, error) {
	// 1. Validate the input data
	// 2. Hash the password
	// 3. Check if the username or email already exists
	// 4. Store user data in a database
	// 5. Generate proper JWT tokens

	pwHash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Error("Could not hash password", "err", err.Error())
		return nil, err
	}

	m := &model.Accounts{}
	m.Id = uuid.New().String()
	m.Email = in.Email
	m.Username = in.Username
	m.Password = string(pwHash)
	if in.Phone != nil {
		m.Phone = *in.Phone
	}

	_, err = l.model.Insert(context.Background(), m)
	if err != nil {
		l.Logger.Error("Could not insert user", "err", err.Error())
		return nil, err
	}

	l.Logger.Info("Generating access token")
	extraClaims := map[string]interface{}{
		"username": in.Username,
		"email":    in.Email,
	}

	access, refresh, err := l.tokens.GenTokenPair(l.ctx, m.Id, extraClaims)
	if err != nil {
		l.Logger.Error("Could not generate access token", "err", err.Error())
		return nil, err
	}

	return &accounts.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
