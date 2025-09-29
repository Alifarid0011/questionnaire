package provider

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func QuizController(service service.QuizService) controller.QuizController {
	return controller.NewQuizController(service)
}
func QuizService(repo repository.QuizRepository) service.QuizService {
	return service.NewQuizService(repo)
}

func QuizRepository(db *mongo.Database) repository.QuizRepository {
	return repository.NewQuizRepository(db)
}
