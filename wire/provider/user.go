package provider

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserService(userRepo repository.UserRepository, casbinRepo repository.CasbinRepository) service.UserService {
	return service.NewUserService(userRepo, casbinRepo)
}
func UserController(userService service.UserService, casbinService service.CasbinService) controller.UserController {
	return controller.NewUserController(userService, casbinService)
}
func UserRepository(db *mongo.Database) repository.UserRepository {
	return repository.NewUserRepository(db)
}
