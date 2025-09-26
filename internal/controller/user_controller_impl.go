package controller

import (
	"errors"
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"github.com/Alifarid0011/questionnaire-back-end/utils/pagination"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type userControllerImpl struct {
	userService   service.UserService
	CasbinService service.CasbinService
}

func NewUserController(userService service.UserService, casbinService service.CasbinService) UserController {
	return &userControllerImpl{userService: userService, CasbinService: casbinService}
}

// FindByUsername
// @Summary      Find user by username
// @Description  Get a user object by username
// @Tags         users
// @Security AuthBearer
// @Param        username  path  string  true  "Username"
// @Success      200 {object} response.Response
// @Router       /users/username/{username} [get]
func (u userControllerImpl) FindByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	userResponse, err := u.userService.FindByUsername(username, ctx.Request.Context())
	if err != nil {
		response.New(ctx).Status(http.StatusNotFound).Errors(errors.New("failed to find user")).MessageID("users.find.failed").Dispatch()
		return
	}
	response.New(ctx).Data(userResponse).Status(http.StatusOK).Message("Information received successfully.").MessageID("users.get.success").Dispatch()
}

// FindUserByUID
// @Summary      Find user by UID
// @Description  Get a user by their UID
// @Tags         users
// @Security AuthBearer
// @Param        uid  path  string  true  "User UID"
// @Success      200 {object} response.Response
// @Router       /users/uid/{uid} [get]
func (u userControllerImpl) FindUserByUID(ctx *gin.Context) {
	uid := ctx.Param("uid")
	objectID, err := primitive.ObjectIDFromHex(uid)
	userResponse, err := u.userService.FindByUID(objectID, ctx.Request.Context())
	if err != nil {
		response.New(ctx).Status(http.StatusNotFound).Errors(errors.New("failed to find user")).MessageID("users.find.failed").Dispatch()
		return
	}
	response.New(ctx).Data(userResponse).Status(http.StatusOK).Message("Information received successfully.").MessageID("users.get.success").Dispatch()
}

// CreateUser
// @Summary      Create a new user
// @Description  Register a new user in the system
// @Tags         users
// @Security AuthBearer
// @Accept       json
// @Produce      json
// @Param        data  body  dto.CreateUserRequest  true  "User data"
// @Success      201  {object}  dto.UserResponse
// @Success      200 {object} response.Response
// @Router       /users [post]
func (u userControllerImpl) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ValidationErrors := utils.GetValidationErrors(err)
		response.New(ctx).Status(http.StatusBadRequest).Errors(errors.New("failed to create user")).Data(map[string]interface{}{
			"errors": ValidationErrors,
		}).MessageID("users.create.failed").Dispatch()
		return
	}
	userResponse, err := u.userService.CreateUser(req, ctx.Request.Context())
	if err != nil {
		response.New(ctx).Status(http.StatusBadRequest).Errors(errors.New("failed to create user")).MessageID("users.create.failed").Dispatch()
		return
	}
	response.New(ctx).Data(userResponse).Status(http.StatusCreated).Message("user created successfully").MessageID("users.create.success").Dispatch()
}

// GetAllUsers
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         users
// @Security     AuthBearer
// @Accept       json
// @Produce      json
// @Param type           query string  false "Pagination type: 'page' or 'cursor'" Enums(page, cursor) default(page)
// @Param page           query int     false "Page number (used with type=page)"
// @Param per_page       query int     false "Items per page (1 to 100)"
// @Param last_seen_id   query string  false "Last seen ID (used with type=cursor)"
// @Param asc            query boolean false "Sort order: true=ASC, false=DESC"
// @Param sort_field     query string  false "Field to sort by" default(_id)
// @Success      200 {object} response.Response
// @Router       /users [get]
func (u userControllerImpl) GetAllUsers(ctx *gin.Context) {
	users, err := u.userService.GetAll(ctx.Request.Context())
	if err != nil {
		response.New(ctx).Errors(err).MessageID("users.get.failed").Status(http.StatusBadRequest).Dispatch()
		return
	}
	paginator, ok := pagination.FromContext(ctx.Request.Context())
	if ok {
		response.New(ctx).Pagination(paginator).Data(users).Status(http.StatusOK).Message("Information received successfully.").MessageID("users.get.success").Dispatch()
		return
	}
	response.New(ctx).Errors(errors.New("failed to retrieve information")).MessageID("users.get.failed").Dispatch()
}

// UpdateUser
// @Summary      Update a user
// @Description  Update user details by UID (admin or self)
// @Tags         users
// @Security AuthBearer
// @Accept       json
// @Produce      json
// @Param        uid   path  string  true  "User UID"
// @Param        data  body  dto.UpdateUserRequest  true  "User update data"
// @Success      200  {object}  dto.UserResponse
// @Router       /users/{uid} [put]
func (u userControllerImpl) UpdateUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ValidationErrors := utils.GetValidationErrors(err)
		response.New(ctx).Status(http.StatusBadRequest).Errors(errors.New("failed to update user")).Data(map[string]interface{}{
			"errors": ValidationErrors,
		}).MessageID("users.create.failed").Dispatch()
		return
	}
	objectID, err := primitive.ObjectIDFromHex(uid)
	userResponse, err := u.userService.UpdateUser(objectID, req, ctx.Request.Context())
	if err != nil {
		response.New(ctx).Status(http.StatusBadRequest).Errors(errors.New("failed to update user")).MessageID("users.update.failed").Dispatch()
		return
	}
	response.New(ctx).Data(userResponse).Status(http.StatusOK).Message("Information received successfully.").MessageID("users.update.success").Dispatch()
}

// DeleteUser
// @Summary      Delete a user
// @Description  Delete user by UID (admin only)
// @Tags         users
// @Security AuthBearer
// @Param        uid path  string  true  "User UID"
// @Success      200 {object} response.Response
// @Router       /users/{uid} [delete]
func (u userControllerImpl) DeleteUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	objectID, _ := primitive.ObjectIDFromHex(uid)
	errUserService := u.userService.DeleteUser(objectID, ctx.Request.Context())
	if errUserService != nil {
		response.New(ctx).Status(http.StatusBadRequest).Errors(errors.New("failed to delete user")).MessageID("users.delete.failed").Dispatch()
		return
	}
	response.New(ctx).Status(http.StatusOK).Message("User deleted successfully").MessageID("users.delete.success").Dispatch()
}

// Me
// @Summary      Get current user info
// @Description  Get authenticated user information from token
// @Tags         users
// @Security AuthBearer
// @Success      200  {object}  dto.UserResponse
// @Router       /users/me [get]
func (u userControllerImpl) Me(ctx *gin.Context) {
	// Retrieve current user ID from the JWT or session context
	userID := ctx.GetString(constant.UserUid)
	permissionsData, errCasbinService := u.CasbinService.GetUserPermissionTree(userID)
	if errCasbinService != nil {
		response.New(ctx).Message("Error retrieving permissions").
			MessageID("casbin.permissions.fetch.failed").
			Status(http.StatusInternalServerError).
			Errors(errCasbinService).
			Dispatch()
		return
	}
	objectID, err := primitive.ObjectIDFromHex(userID)
	userResponse, err := u.userService.Me(objectID, ctx.Request.Context())
	if err != nil {
		response.New(ctx).Errors(err).Message("Operation failed").MessageID("users.me.get.failed").Status(http.StatusInternalServerError).Dispatch()
		return
	}
	response.New(ctx).Message("Information received successfully.").MessageID("users.me.get.success").Data(map[string]interface{}{
		"user":        userResponse,
		"permissions": permissionsData,
	}).Status(http.StatusOK).Dispatch()
}
