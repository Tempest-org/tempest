package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OrganizationsInvitationsModel = (*customOrganizationsInvitationsModel)(nil)

type (
	// OrganizationsInvitationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrganizationsInvitationsModel.
	OrganizationsInvitationsModel interface {
		organizationsInvitationsModel
		withSession(session sqlx.Session) OrganizationsInvitationsModel
	}

	customOrganizationsInvitationsModel struct {
		*defaultOrganizationsInvitationsModel
	}
)

// NewOrganizationsInvitationsModel returns a model for the database table.
func NewOrganizationsInvitationsModel(conn sqlx.SqlConn) OrganizationsInvitationsModel {
	return &customOrganizationsInvitationsModel{
		defaultOrganizationsInvitationsModel: newOrganizationsInvitationsModel(conn),
	}
}

func (m *customOrganizationsInvitationsModel) withSession(session sqlx.Session) OrganizationsInvitationsModel {
	return NewOrganizationsInvitationsModel(sqlx.NewSqlConnFromSession(session))
}
