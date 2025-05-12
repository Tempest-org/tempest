package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OrganizationsMembersModel = (*customOrganizationsMembersModel)(nil)

type (
	// OrganizationsMembersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrganizationsMembersModel.
	OrganizationsMembersModel interface {
		organizationsMembersModel
		withSession(session sqlx.Session) OrganizationsMembersModel
	}

	customOrganizationsMembersModel struct {
		*defaultOrganizationsMembersModel
	}
)

// NewOrganizationsMembersModel returns a model for the database table.
func NewOrganizationsMembersModel(conn sqlx.SqlConn) OrganizationsMembersModel {
	return &customOrganizationsMembersModel{
		defaultOrganizationsMembersModel: newOrganizationsMembersModel(conn),
	}
}

func (m *customOrganizationsMembersModel) withSession(session sqlx.Session) OrganizationsMembersModel {
	return NewOrganizationsMembersModel(sqlx.NewSqlConnFromSession(session))
}
