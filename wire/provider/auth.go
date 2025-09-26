package provider

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/controller"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/internal/service"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthService(
	userRepo repository.UserRepository,
	refreshTokenRepo repository.RefreshTokenRepository,
	tokenManager utils.JwtToken,
	blackListRepo repository.BlackListTokenRepository,
	casbinRepo repository.CasbinRepository,
) service.AuthService {
	return service.NewAuthService(userRepo, tokenManager, refreshTokenRepo, blackListRepo, casbinRepo)
}

func AuthController(authService service.AuthService, userService service.UserService) controller.AuthController {
	return controller.NewAuthController(authService, userService)
}
func JWT(secret string) utils.JwtToken {
	return utils.NewJwtToken(secret)
}
func BlackListRepository(db *mongo.Database) repository.BlackListTokenRepository {
	return repository.NewBlackListRepository(db)
}
func RefreshTokenRepository(db *mongo.Database) repository.RefreshTokenRepository {
	return repository.NewRefreshTokenRepository(db)
}

func RouterController(engine *gin.Engine) controller.RouteController {
	return controller.NewRouteController(engine)
}
