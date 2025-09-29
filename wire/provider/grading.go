package provider

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
)

func GradingService(
	quizRepo repository.QuizRepository,
	userAnswerRepo repository.UserAnswerRepository,
	apiClient utils.ShortAnswerAPIClient,
) service.GradingService {
	return service.NewGradingService(quizRepo, userAnswerRepo, apiClient)
}

// UserAnswerGradingController returns a new UserAnswerGradingController instance
func UserAnswerGradingController(gs service.GradingService) controller.UserAnswerGradingController {
	return controller.NewUserAnswerGradingController(gs)
}

func ProvideShortAnswerAPIClient() utils.ShortAnswerAPIClient {
	return utils.NewShortAnswerAPIClient()
}
