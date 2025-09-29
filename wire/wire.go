//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"github.com/Alifarid0011/questionnaire-back-end/wire/provider"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	// Core Components
	TokenManager  utils.JwtToken
	BlackListRepo repository.BlackListTokenRepository
	RouterCtr     controller.RouteController
	// Auth
	AuthCtrl         controller.AuthController
	RefreshTokenRepo repository.RefreshTokenRepository
	Enforcer         *casbin.Enforcer
	Mongo            *mongo.Client
	Engine           *gin.Engine
	// Casbin/ACL
	CasbinRepo    repository.CasbinRepository
	CasbinCtrl    controller.CasbinController
	CasbinService service.CasbinService
	// User
	UserCtrl    controller.UserController
	UserService service.UserService
	UserRepo    repository.UserRepository
	//Quiz
	QuizCtrl    controller.QuizController
	QuizService service.QuizService
	QuizRepo    repository.QuizRepository
	//	user Answer
	UserAnswerCtrl    controller.UserAnswerController
	UserAnswerService service.UserAnswerService
	UserAnswerRepo    repository.UserAnswerRepository
	//	Comment
	CommentCtrl    controller.CommentController
	CommentService service.CommentService
	CommentRepo    repository.CommentRepository
	// question rating
	QuestionRatingCtrl    controller.QuestionRatingController
	QuestionRatingService service.QuestionRatingService
	QuestionRatingRepo    repository.QuestionRatingRepository
	//	 grading
	GradingCtrl       controller.UserAnswerGradingController
	GradingService    service.GradingService
	ShortAnswerClient utils.ShortAnswerAPIClient
}

// InitializeApp initializes the application with all its dependencies.
func InitializeApp(secret string) (*App, error) {
	wire.Build(
		//Mongo
		provider.MongoClient,
		provider.Database,
		provider.RouterEngine,
		// Casbin/ACL
		provider.CasbinController,
		provider.CasbinService,
		provider.CasbinRepository,
		provider.CasbinEnforcer,
		// User
		provider.UserController,
		provider.UserService,
		provider.UserRepository,
		// Auth
		provider.AuthController,
		provider.AuthService,
		provider.JWT,
		provider.BlackListRepository,
		provider.RefreshTokenRepository,
		provider.RouterController,
		// Quiz
		provider.QuizRepository,
		provider.QuizService,
		provider.QuizController,
		//User Answer
		provider.UserAnswerController,
		provider.UserAnswerRepository,
		provider.UserAnswerService,
		// comment
		provider.CommentController,
		provider.CommentRepository,
		provider.CommentService,
		// question rating
		provider.QuestionRatingRepository,
		provider.QuestionRatingService,
		provider.QuestionRatingController,
		//	 grading
		provider.UserAnswerGradingController,
		provider.ProvideShortAnswerAPIClient,
		provider.GradingService,
		wire.Struct(new(App),
			"*"))
	return &App{}, nil
}
