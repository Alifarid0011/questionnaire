//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Alifarid0011/questionnaire-back-end/wire/provider"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type App struct {
	Engine *gin.Engine
}

// InitializeApp initializes the application with all its dependencies.
func InitializeApp() (*App, error) {
	wire.Build(
		provider.ProvideRouterEngine,
		wire.Struct(new(App),
			"*"))
	return &App{}, nil
}
