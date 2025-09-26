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
	return r
}
