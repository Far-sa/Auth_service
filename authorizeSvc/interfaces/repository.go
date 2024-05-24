package interfaces

type RoleRepository interface {
	AssignRole(username, role string) error
	CheckPermission(username, permission string) (bool, error)
}
