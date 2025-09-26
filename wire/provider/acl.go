package provider

import (
	"github.com/Alifarid0011/questionnaire-back-end/config"
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/casbin/casbin/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func CasbinEnforcer(mongoClient *mongo.Client) *casbin.Enforcer {
	return config.InitCasbin(mongoClient)
}
func CasbinController(service service.CasbinService) controller.CasbinController {
	return controller.NewACLController(service)
}
func CasbinService(repo repository.CasbinRepository) service.CasbinService {
	return service.NewCasbinService(repo)
}

func CasbinRepository(enforcer *casbin.Enforcer, client *mongo.Client) repository.CasbinRepository {
	return repository.NewCasbinRepository(enforcer, client.Database("casbin"))
}
