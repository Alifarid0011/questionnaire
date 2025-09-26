package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func RegisterAclRoutes(r *gin.Engine, app *wire.App) {
	aclRouter := r.Group("/acl")
	{
		aclRouter.GET("/check", app.CasbinCtrl.CheckPermission)
		aclRouter.GET("/permissions", app.CasbinCtrl.ListAllCasbinData)
		aclRouter.GET("/roles", app.CasbinCtrl.Roles)
		aclRouter.GET("/user_roles", app.CasbinCtrl.UserRoles)
		//policies
		policiesRouter := aclRouter.Group("/policies")
		{
			policiesRouter.POST("", app.CasbinCtrl.CreatePolicy)
			policiesRouter.GET("/:sub/permissions", app.CasbinCtrl.PermissionsTree)
			policiesRouter.DELETE("", app.CasbinCtrl.RemovePolicy)

		}
		//policy group
		policyGroupRouter := aclRouter.Group("/policy_group")
		{
			policyGroupRouter.POST("", app.CasbinCtrl.AddGroupingPolicy)
			policyGroupRouter.DELETE("", app.CasbinCtrl.RemoveGroupingPolicy)
		}
	}
}
