package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/middleware"
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *wire.App) *gin.Engine {
	r := app.Engine
	r.Use(middleware.CORSMiddleware())
	RegisterSwaggerRoutes(r)
	RegisterListRoutes(r, app)
	RegisterQuizRoutes(r, app)
	RegisterAclRoutes(r, app)
	RegisterUserRoutes(r, app)
	RegisterAuthRoutes(r, app)
	return r
}
