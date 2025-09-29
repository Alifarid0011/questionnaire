package provider

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func QuestionRatingController(service service.QuestionRatingService) controller.QuestionRatingController {
	return controller.NewQuestionRatingController(service)
}
func QuestionRatingService(repo repository.QuestionRatingRepository) service.QuestionRatingService {
	return service.NewQuestionRatingService(repo)
}

func QuestionRatingRepository(db *mongo.Database) repository.QuestionRatingRepository {
	return repository.NewQuestionRatingRepository(db)
}
