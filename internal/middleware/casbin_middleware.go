package middleware

import (
	"errors"
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := enforcer.GetRolesForUser(c.GetString("user_uid"))
		if err != nil {
			response.New(c).Message("Operation failed.").
				MessageID("role.fetch.error").
				Status(http.StatusConflict).
				Errors(err).
				Dispatch()
			return
		}
		act := c.Request.Method
		obj := c.Request.URL.Path
		user := c.GetString("user_uid")
		ok, errEnforce := enforcer.Enforce(user, obj, act)
		if errEnforce != nil {
			response.New(c).Message("Access check encountered an error.").
				MessageID("casbin.enforce.error").
				Status(http.StatusForbidden).
				Errors(errEnforce).
				Dispatch()
			return
		}
		if !ok {
			response.New(c).Message("You do not have the required permissions.").
				MessageID("casbin.enforce.Unauthorized").
				Status(http.StatusForbidden).
				Errors(errors.New("required permissions are missing")).
				Dispatch()
			return
		}
		c.Set(constant.ContextRolesKey, roles)
		c.Next()
	}
}
