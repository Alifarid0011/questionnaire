package routers

import (
	"github.com/Alifarid0011/questionnaire-back-end/wire"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *wire.App) *gin.Engine {
	r := app.Engine
	//Todo: registration routers
	return r
}
