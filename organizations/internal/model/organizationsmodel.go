package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OrganizationsModel = (*customOrganizationsModel)(nil)

type (
	// OrganizationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrganizationsModel.
	OrganizationsModel interface {
		organizationsModel
		withSession(session sqlx.Session) OrganizationsModel
	}

	customOrganizationsModel struct {
		*defaultOrganizationsModel
	}
)

// NewOrganizationsModel returns a model for the database table.
func NewOrganizationsModel(conn sqlx.SqlConn) OrganizationsModel {
	return &customOrganizationsModel{
		defaultOrganizationsModel: newOrganizationsModel(conn),
	}
}

func (m *customOrganizationsModel) withSession(session sqlx.Session) OrganizationsModel {
	return NewOrganizationsModel(sqlx.NewSqlConnFromSession(session))
}
