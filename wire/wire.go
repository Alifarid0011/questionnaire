//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/Alifarid0011/questionnaire-back-end/wire/provider"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Enforcer *casbin.Enforcer
	Mongo    *mongo.Client
	Engine   *gin.Engine
	// Casbin/ACL
	CasbinRepo    repository.CasbinRepository
	CasbinCtrl    controller.CasbinController
	CasbinService service.CasbinService
}

// InitializeApp initializes the application with all its dependencies.
func InitializeApp() (*App, error) {
	wire.Build(
		provider.RouterEngine,
		// Casbin/ACL
		provider.CasbinController,
		provider.CasbinService,
		provider.CasbinRepository,
		provider.MongoClient,
		provider.CasbinEnforcer,
		wire.Struct(new(App),
			"*"))
	return &App{}, nil
}
