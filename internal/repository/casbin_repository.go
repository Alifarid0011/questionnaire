package repository

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
)

type CasbinRepository interface {
	Enforce(sub, obj, act, AllowOrDeny string) (bool, error)
	AddPolicy(sub, obj, act, AllowOrDeny string) (bool, error)
	RemovePolicy(sub, obj, act, AllowOrDeny string) (bool, error)
	GetPolicies() ([]models.CasbinPolicy, error)
	AddGroupingPolicy(child, parent string) (bool, error)
	RemoveGroupingPolicy(child, parent string) (bool, error)
	GetGroupingPolicies() ([][]string, error)
	GetRolesForUser(userUID string) ([]string, error)
	GetPermissionsGroupedBySubject() ([]models.SubjectWithPermissions, error)
	GetPermissionsBySubjects(subjects []string) ([]models.PermissionDTO, error)
	GetAllRolesFromPolicies(ctx context.Context) ([]string, error)
	GetAllPoliciesRelatedToUser(userUID string) ([]models.PermissionDTO, error, []string)
}
