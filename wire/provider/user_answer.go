package provider

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserAnswerController(service service.UserAnswerService) controller.UserAnswerController {
	return controller.NewUserAnswerController(service)
}
func UserAnswerService(repo repository.UserAnswerRepository) service.UserAnswerService {
	return service.NewUserAnswerService(repo)
}

func UserAnswerRepository(db *mongo.Database) repository.UserAnswerRepository {
	return repository.NewUserAnswerRepository(db)
}
