package authorization

import (
	"context"
	"fmt"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	"iceline-hosting.com/core/dbhelper"
)

const (
	// TODO: review the model file path
	modelFile = `./basic_model.csv`
)

type CoreAuthorizer struct {
	enforcer *casbin.Enforcer
	db       *sqlx.DB
}

func NewCoreAuthorizer(sdb *sqlx.DB, tableName string) (*CoreAuthorizer, error) {
	a, err := sqladapter.NewAdapter(sdb.DB, "postgres", tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to setup adapter. %w", err)
	}

	e, err := casbin.NewEnforcer(modelFile, a)
	if err != nil {
		return nil, fmt.Errorf("failed to setup enforcer. %w", err)
	}

	return &CoreAuthorizer{
		enforcer: e,
		db:       sdb,
	}, nil
}

func (a *CoreAuthorizer) IsAuthorized(subject, endpoint, action string) (bool, error) {
	return a.enforcer.Enforce(subject, endpoint, action)
}

func (a *CoreAuthorizer) GetUserPermissions(userID string) (map[string][]string, error) {
	userroles, err := dbhelper.GetUserPermissions(context.Background(), a.db, userID)
	if err != nil {
		return nil, err
	}

	permissions := make(map[string][]string)

	for _, userrole := range userroles {
		if _, exists := permissions[userrole.Role]; !exists {
			permissions[userrole.Role] = []string{userrole.Resource}
		} else {
			permissions[userrole.Role] = append(permissions[userrole.Role], userrole.Resource)
		}
	}

	return permissions, nil
}

func (a *CoreAuthorizer) SetUserPermission(userID, role, resource string) error {
	return dbhelper.SetUserPermission(context.Background(), a.db, userID, role, resource)
}

func (a *CoreAuthorizer) AddPolicy(subject, endpoint, action string) error {
	_, err := a.enforcer.AddPolicy(subject, endpoint, action)
	if err != nil {
		return err
	}

	return nil
}

func (a *CoreAuthorizer) AddGroupingPolicy(subject, role string) error {
	_, err := a.enforcer.AddGroupingPolicy(subject, role)
	if err != nil {
		return err
	}

	return nil
}
