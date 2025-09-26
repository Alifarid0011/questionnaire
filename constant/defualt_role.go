package constant

type UserRole = string

const (
	RoleSuperAdmin UserRole = "super_admin"
	RoleUser       UserRole = "user"
)

var DefaultRoles = []string{RoleSuperAdmin, RoleUser}
var DefaultPermissions = []struct {
	Object      string
	Action      string
	Attribute   string
	AllowOrDeny string
	Subject     string
	Entity      string
}{
	{Object: "/users/me", Action: "GET", AllowOrDeny: "allow", Subject: RoleUser},
	{Object: "*", Action: "*", AllowOrDeny: "allow", Subject: RoleSuperAdmin},
}
