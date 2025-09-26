package controller

import (
	"errors"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ACLController struct {
	service service.CasbinService
}

func NewACLController(srv service.CasbinService) *ACLController {
	return &ACLController{service: srv}
}

// CheckPermission checks user permission level.
// @Summary Check Permission
// @Description Checks whether the user is allowed to perform a specific action on a resource.
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param request body dto.CheckPermissionDTO true "Permission check information"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/check [get]
func (ctl *ACLController) CheckPermission(ctx *gin.Context) {
	var req dto.CheckPermissionDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("Invalid request data.").
			MessageID("acl.check.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	allowed, err := ctl.service.IsAllowed(req.Sub, req.Obj, req.Act, req.AllowOrDeny)
	if err != nil {
		response.New(ctx).Message("Permission check failed.").
			MessageID("acl.check.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("Permission check completed.").
		MessageID("acl.check.success").
		Status(http.StatusOK).
		Data(gin.H{"allowed": allowed}).
		Dispatch()
}

// CreatePolicy defines a new permission policy.
// @Summary Add Policy
// @Description Adds a policy for a user/group to perform an action on a resource.
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param request body dto.CheckPermissionDTO true "Policy information"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/policies [post]
func (ctl *ACLController) CreatePolicy(ctx *gin.Context) {
	var req dto.CheckPermissionDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("Invalid request data.").
			MessageID("acl.policy.create.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	added, err := ctl.service.GrantPermission(req.Sub, req.Obj, req.Act, req.AllowOrDeny)
	if err != nil || !added {
		response.New(ctx).Message("Failed to register policy.").
			MessageID("acl.policy.create.failed").
			Status(http.StatusInternalServerError).
			Errors(errors.New("failed to register policy")).
			Dispatch()
		return
	}
	response.New(ctx).Message("Policy successfully registered.").
		MessageID("acl.policy.create.success").
		Status(http.StatusCreated).
		Dispatch()
}

// ListAllCasbinData retrieves all policies and groupings from Casbin.
// @Summary Get All Casbin Data
// @Description Retrieves all policies (p) and grouping policies (g, g2) stored in Casbin.
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/permissions [get]
func (ctl *ACLController) ListAllCasbinData(ctx *gin.Context) {
	data, err := ctl.service.GetAllCasbinData()
	if err != nil {
		response.New(ctx).Message("Error retrieving Casbin data.").
			MessageID("casbin.data.fetch.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}

	response.New(ctx).Message("All policies and groupings successfully retrieved.").
		MessageID("casbin.data.fetch.success").
		Status(http.StatusOK).
		Data(data).
		Dispatch()
}

// RemovePolicy removes a specific policy.
// @Summary Remove Policy
// @Description Removes the specified permission policy for the user/group.
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param request body dto.CheckPermissionDTO true "Policy information for removal"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/policies [delete]
func (ctl *ACLController) RemovePolicy(ctx *gin.Context) {
	var req dto.CheckPermissionDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("Invalid request data.").
			MessageID("acl.policy.remove.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	removed, err := ctl.service.RevokePermission(req.Sub, req.Obj, req.Act, req.AllowOrDeny)
	if err != nil || !removed {
		response.New(ctx).Message("Policy removal failed.").
			MessageID("acl.policy.remove.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("Policy successfully removed.").
		MessageID("acl.policy.remove.success").
		Status(http.StatusOK).
		Dispatch()
}

// AddGroupingPolicy adds a grouping policy (g).
// @Summary Add Grouping Policy (g)
// @Description Assigns one role or group to another role/group (g grouping).
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param request body dto.GroupingDTO true "Grouping information"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/policy_group [post]
func (ctl *ACLController) AddGroupingPolicy(ctx *gin.Context) {
	var req dto.GroupingDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("Invalid request data.").
			MessageID("acl.grouping.add.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	err := ctl.service.AddGrouping(req.Parent, req.Child)
	if err != nil {
		response.New(ctx).Message("Adding grouping policy failed.").
			MessageID("acl.grouping.add.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("Grouping successfully added.").
		MessageID("acl.grouping.add.success").
		Status(http.StatusOK).
		Dispatch()
}

// RemoveGroupingPolicy removes a grouping policy (g).
// @Summary Remove Grouping Policy (g)
// @Description Removes the link between one role/group and another role/group.
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param request body dto.GroupingDTO true "Grouping information for removal"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/policy_group [delete]
func (ctl *ACLController) RemoveGroupingPolicy(ctx *gin.Context) {
	var req dto.GroupingDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("Invalid request data.").
			MessageID("acl.grouping.remove.invalid").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	err := ctl.service.RemoveGrouping(req.Parent, req.Child)
	if err != nil {
		response.New(ctx).Message("Grouping removal failed.").
			MessageID("acl.grouping.remove.failed").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("Grouping successfully removed.").
		MessageID("acl.grouping.remove.success").
		Status(http.StatusOK).
		Dispatch()
}

// PermissionsTree retrieves the permission tree for a subject.
// @Summary Get Subject Permission Tree
// @Description Retrieves hierarchical permissions based on user ID or role (subject).
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param sub path string true "User ID or Role (subject)"
// @Success 200 {object} response.Response
// @Failure 400,500 {object} response.Response
// @Router /acl/policies/{sub}/permissions [get]
func (ctl *ACLController) PermissionsTree(ctx *gin.Context) {
	sub := ctx.Param("sub")
	data, err := ctl.service.GetUserPermissionTree(sub)
	if err != nil {
		response.New(ctx).Message("Error retrieving permissions.").
			MessageID("casbin.permissions.fetch.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("Permissions successfully retrieved.").
		MessageID("casbin.permissions.fetch.success").
		Status(http.StatusOK).
		Data(data).
		Dispatch()
}

// Roles retrieves a list of all defined roles from policies.
// @Summary Get Roles List
// @Description This endpoint returns all roles defined as v0 in policies (excluding those that are ObjectID).
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /acl/roles [get]
func (ctl *ACLController) Roles(ctx *gin.Context) {
	roles, err := ctl.service.GetAllRoles(ctx)
	if err != nil {
		response.New(ctx).
			Message("Error retrieving roles.").
			MessageID("casbin.roles.fetch.error").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).
		Message("Roles successfully retrieved.").
		MessageID("casbin.roles.fetch.success").
		Status(http.StatusOK).
		Data(dto.RolesResponse{Roles: roles}).
		Dispatch()
}

// UserRoles retrieves roles assigned to a user.
// @Summary Get User Roles
// @Description Retrieves all roles assigned to the given user ID.
// @Tags ACL
// @Security AuthBearer
// @Accept json
// @Produce json
// @Param uid query string true "user_uid" example(64b2fa75e7d1f4a739fa2a11)
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /acl/user_roles [get]
func (ctl *ACLController) UserRoles(ctx *gin.Context) {
	var req dto.UIDQuery
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.New(ctx).Message("Invalid request data.").
			MessageID("acl.query.invalid"). // Changed to a more general error ID
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	roles, err := ctl.service.GetRolesForUser(req.UID)
	if err != nil {
		response.New(ctx).
			Message("Error retrieving user roles.").
			MessageID("casbin.user.roles.fetch.error"). // Specific ID
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).
		Message("User roles successfully retrieved.").
		MessageID("casbin.user.roles.fetch.success"). // Specific ID
		Status(http.StatusOK).
		Data(dto.RolesResponse{Roles: roles}).
		Dispatch()
}
