package authorization

type Authorizer interface {
	IsAuthorized(subject, endpoint, action string) (bool, error)

	// GetUserPermissions should return in format of `role => [resource1, resource2]`
	GetUserPermissions(userID string) (map[string][]string, error)

	SetUserPermission(userID, role, resource string) error

	AddPolicy(subject, endpoint, action string) error
}

const (
	DefaultRole = "users"
	OwnerRole   = "owner"
)
