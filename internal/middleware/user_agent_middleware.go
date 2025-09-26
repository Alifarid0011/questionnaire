package middleware

import (
	"github.com/Alifarid0011/questionnaire-back-end/constant"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"github.com/gin-gonic/gin"
)

func UserAgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ua := new(utils.UserAgent).Constructor(c)
		c.Set(constant.UserAgentInfo, ua)
		c.Next()
	}
}
