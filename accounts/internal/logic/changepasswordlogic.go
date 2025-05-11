package logic

import (
	"context"

	"accounts/accounts"
	"accounts/internal/model"
	"accounts/internal/svc"
	"accounts/pkg/datasource"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	vl *ValidateLogic
	ll *LoginLogic
	model model.AccountsModel
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	logger := logx.WithContext(ctx)
	vl := NewValidateLogic(ctx, svcCtx)
	ll:= NewLoginLogic(ctx, svcCtx)
	pgConn := datasource.NewPostgresConn(svcCtx.Config.Database.URI)
	accModel := model.NewAccountsModel(pgConn)

	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger,
		vl:     vl,
		ll:     ll,
		model:  accModel,
	}
}

func (l *ChangePasswordLogic) ChangePassword(in *accounts.ChangePasswordRequest) (*accounts.ChangePasswordResponse, error) {
	acc, err := l.ll.accountByIdentifier(in.Identifier, in.Value)
	if err != nil {
		l.Logger.Error("Could not find account", "err", err.Error())
		return nil, errorx.Wrap(err, "Could not find Account")
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(in.OldPassword))
	if 	err != nil {
		l.Logger.Error("Invalid old password", "err", err.Error())
		return nil, errorx.Wrap(err, "Invalid old password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Error("Could not hash new password", "err", err.Error())
		return nil, errorx.Wrap(err, "Could not hash new password")
	}

	acc.Password = string(hash)
	err = l.model.Update(l.ctx, acc)
	if err != nil {
		l.Logger.Error("Could not update account", "err", err.Error())
		return nil, errorx.Wrap(err, "Could not update account")
	}

	return &accounts.ChangePasswordResponse{
		Success: true,
		Message: "OK",
	}, nil
}
