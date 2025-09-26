package service

import (
	"context"
	"fmt"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"strings"
)

type CasbinServiceImpl struct {
	repo repository.CasbinRepository
}

func NewCasbinService(repo repository.CasbinRepository) CasbinService {
	return &CasbinServiceImpl{repo: repo}
}

func (s *CasbinServiceImpl) AddGrouping(parent string, child string) error {
	added, err := s.repo.AddGroupingPolicy(child, parent)
	if err != nil {
		return err
	}
	if !added {
		return fmt.Errorf("already exist")
	}
	return nil
}

func (s *CasbinServiceImpl) RemoveGrouping(parent string, child string) error {
	removed, err := s.repo.RemoveGroupingPolicy(child, parent)
	if err != nil {
		return err
	}
	if !removed {
		return fmt.Errorf("dose not exist")
	}
	return nil
}

func (s *CasbinServiceImpl) IsAllowed(sub, act, obj, attr, AllowOrDeny, entity string) (bool, error) {
	return s.repo.Enforce(sub, obj, act, attr, AllowOrDeny, entity)
}

func (s *CasbinServiceImpl) GrantPermission(sub, obj, act, attr, AllowOrDeny, entity string) (bool, error) {
	return s.repo.AddPolicy(sub, obj, act, attr, AllowOrDeny, entity)
}

func (s *CasbinServiceImpl) RevokePermission(sub, obj, act, attr, AllowOrDeny, entity string) (bool, error) {
	return s.repo.RemovePolicy(sub, obj, act, attr, AllowOrDeny, entity)
}
func (s *CasbinServiceImpl) GetAllCasbinData() (map[string]interface{}, error) {
	policies, err := s.repo.GetPolicies()
	if err != nil {
		return nil, err
	}
	groupingPoliciesG, err := s.repo.GetGroupingPolicies()
	return map[string]interface{}{
		"policies":            policies,
		"grouping_policies_g": groupingPoliciesG,
	}, nil
}
func (s *CasbinServiceImpl) ListPermissions() ([]models.CasbinPolicy, error) {
	return s.repo.GetPolicies()
}

func (s *CasbinServiceImpl) GetPermissionsBySubject() ([]models.SubjectWithPermissions, error) {
	return s.repo.GetPermissionsGroupedBySubject()
}
func (s *CasbinServiceImpl) GetAllRoles(ctx context.Context) ([]string, error) {
	roles, err := s.repo.GetAllRolesFromPolicies(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get roles from policies: %w", err)
	}
	return roles, nil
}

type PermissionNode map[string]interface{}

func (s *CasbinServiceImpl) GetRolesForUser(userUID string) ([]string, error) {
	return s.repo.GetRolesForUser(userUID)
}

func (s *CasbinServiceImpl) GetUserPermissionTree(userUID string) (PermissionNode, error) {
	policies, err, roles := s.repo.GetAllPoliciesRelatedToUser(userUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get relevant permissions: %w", err)
	}
	grouped := make(map[string][]models.PermissionDTO)
	for _, p := range policies {
		grouped[p.Sub] = append(grouped[p.Sub], p)
	}
	var roleTrees []PermissionNode
	for _, role := range roles {
		if perms, ok := grouped[role]; ok {
			roleTrees = append(roleTrees, BuildTreePermission(perms))
		}
	}
	userTree := make(PermissionNode)
	if userPerms, ok := grouped[userUID]; ok {
		userTree = BuildTreePermission(userPerms)
	}
	finalTree := MergePermissionTreesWithUserPriority(roleTrees, userTree)
	return finalTree, nil
}
func MergePermissionTreesWithUserPriority(roleTrees []PermissionNode, userTree PermissionNode) PermissionNode {
	merged := make(PermissionNode)

	for _, roleTree := range roleTrees {
		mergeInto(merged, roleTree, false)
	}
	mergeInto(merged, userTree, true)

	return merged
}
func mergeInto(dst, src PermissionNode, canOverride bool) {
	for key, srcVal := range src {
		if key == "actions" {
			if dstActions, ok := dst[key].(map[string]string); ok {
				for act, eft := range srcVal.(map[string]string) {
					_, exists := dstActions[act]
					if !exists || canOverride {
						dstActions[act] = eft
					}
				}
			} else {
				dst[key] = srcVal
			}
			continue
		}
		if _, ok := dst[key]; !ok {
			dst[key] = make(PermissionNode)
		}
		dstChild := dst[key].(PermissionNode)
		srcChild := srcVal.(PermissionNode)
		mergeInto(dstChild, srcChild, canOverride)
	}
}
func BuildTreePermission(perms []models.PermissionDTO) PermissionNode {
	root := make(PermissionNode)
	for _, perm := range perms {
		parts := strings.Split(strings.Trim(perm.Obj, "/"), "/")
		current := root
		for i, part := range parts {
			if _, ok := current[part]; !ok {
				current[part] = make(PermissionNode)
			}
			if i == len(parts)-1 {
				node := current[part].(PermissionNode)
				if _, exists := node["actions"]; !exists {
					node["actions"] = make(map[string]string)
				}
				actions := node["actions"].(map[string]string)
				actions[perm.Act] = perm.Eft
			} else {
				current = current[part].(PermissionNode)
			}
		}
	}

	return root
}
