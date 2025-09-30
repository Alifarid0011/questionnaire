package controller

import (
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authControllerImpl struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthController(authService service.AuthService, userService service.UserService) AuthController {
	return &authControllerImpl{authService: authService, userService: userService}
}

// Login godoc
// @Summary      Authenticate user
// @Description  Takes username and password, returns access and refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login credentials"
// @Success      200 {object} response.Response
// @Router       /auth/login [post]
func (c *authControllerImpl) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("operation failed").
			MessageID("auth.login.failed").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	UserAgentInfo, _ := ctx.Get(constant.UserAgentInfo)
	resp, err := c.authService.Login(req, UserAgentInfo.(*utils.UserAgent))
	if err != nil {
		response.New(ctx).Message("operation failed").
			MessageID("auth.login.failed").
			Status(http.StatusUnauthorized).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("welcome").
		MessageID("auth.login.success").
		Data(resp).
		Status(http.StatusOK).
		Dispatch()
}

// Logout godoc
// @Summary Logout and invalidate refresh token
// @Description Logs out the user and invalidates the refresh token based on user-agent
// @Tags Auth
// @Accept json
// @Produce json
// @Security AuthBearer
// @Success      200 {object} response.Response
// @Router /auth/logout [get]
func (c *authControllerImpl) Logout(ctx *gin.Context) {
	UserAgentInfo, _ := ctx.Get(constant.UserAgentInfo)
	token, _ := ctx.Get("access_token")
	err := c.authService.Logout(token.(string), UserAgentInfo.(*utils.UserAgent))
	if err != nil {
		response.New(ctx).Message("operation failed").
			MessageID("auth.logout.failed").
			Status(http.StatusUnauthorized).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("successfully logout").
		MessageID("auth.logout.success").
		Status(http.StatusOK).
		Data(map[string]string{
			"message": "bye ...",
		}).
		Dispatch()
}

// Register godoc
// @Summary Register  user Registration
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "register user"
// @Success      200 {object} response.Response
// @Router /auth/register [post]
func (c *authControllerImpl) Register(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("operation failed").
			MessageID("auth.register.failed").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	userResponse, err := c.userService.CreateUser(req, ctx.Request.Context())
	if err != nil {
		response.New(ctx).Message("operation failed").
			MessageID("auth.register.failed").
			Status(http.StatusInternalServerError).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("welcome").
		MessageID("auth.register.success").
		Data(userResponse).
		Status(http.StatusCreated).
		Dispatch()
}

// UseRefreshToken godoc
// @Summary Use refresh token to get new access token
// @Description Uses a refresh token and user-agent info to generate a new access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshRequest true "Refresh token payload"
// @Success      200 {object} response.Response
// @Router /auth/refresh_token [post]
func (c *authControllerImpl) UseRefreshToken(ctx *gin.Context) {
	var req dto.RefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.New(ctx).Message("operation failed").
			MessageID("auth.refresh_token.failed").
			Status(http.StatusBadRequest).
			Errors(err).
			Dispatch()
		return
	}
	UserAgentInfo, _ := ctx.Get(constant.UserAgentInfo)
	resp, err := c.authService.UseRefreshToken(req, UserAgentInfo.(*utils.UserAgent))
	if err != nil {
		response.New(ctx).Message("operation failed").
			MessageID("auth.refresh_token.failed").
			Status(http.StatusUnauthorized).
			Errors(err).
			Dispatch()
		return
	}
	response.New(ctx).Message("token generated successfully").
		MessageID("auth.refresh_token.success").
		Data(resp).
		Status(http.StatusOK).
		Dispatch()
}
