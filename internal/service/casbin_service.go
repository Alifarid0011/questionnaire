package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
)

type CasbinService interface {
	IsAllowed(sub, obj, act, AllowOrDeny string) (bool, error)
	GrantPermission(sub, obj, act, AllowOrDeny string) (bool, error)
	RevokePermission(sub, obj, act, AllowOrDeny string) (bool, error)
	ListPermissions() ([]models.CasbinPolicy, error)
	AddGrouping(parent string, child string) error
	GetAllCasbinData() (map[string]interface{}, error)
	RemoveGrouping(parent string, child string) error
	GetPermissionsBySubject() ([]models.SubjectWithPermissions, error)
	GetAllRoles(ctx context.Context) ([]string, error)
	GetRolesForUser(userUID string) ([]string, error)
	GetUserPermissionTree(userUID string) (PermissionNode, error)
}
