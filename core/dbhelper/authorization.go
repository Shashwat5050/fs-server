package dbhelper

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"iceline-hosting.com/core/models"
)

// GetUserPermissions returns the user's permissions.
func GetUserPermissions(ctx context.Context, db *sqlx.DB, userId string) ([]models.UserRole, error) {
	if db == nil {
		return nil, errors.New("database is not initialized")
	}

	var userRoles []models.UserRole
	err := db.SelectContext(ctx, &userRoles, "SELECT ur.id, ur.user_id, ur.role, ur.resource FROM user_roles as ur left join gs_info as gi on ur.resource=gi.server_name WHERE ur.user_id=$1", userId)
	if err != nil {
		return nil, err
	}

	return userRoles, nil
}

// SetUserPermission sets the user's permission to the given role and resource.
func SetUserPermission(ctx context.Context, db *sqlx.DB, userID, role, resource string) error {
	if db == nil {
		return errors.New("database is not initialized")
	}

	permission := models.UserRole{
		UserID:   userID,
		Role:     role,
		Resource: resource,
	}

	_, err := db.NamedExecContext(ctx, "INSERT INTO user_roles (user_id, role, resource) VALUES (:user_id, :role, :resource)", permission)
	return err
}
