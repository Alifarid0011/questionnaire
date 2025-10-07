package middleware

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto/response"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AuthMiddleware(blackRepo repository.BlackListTokenRepository, tokenManager utils.JwtToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString, err := utils.ExtractBearerToken(authHeader)
		if err != nil {
			response.New(c).Status(http.StatusUnauthorized).Errors(err).Message("invalid token").MessageID("auth.middleware.failed").Dispatch()
			return
		}
		token, err := tokenManager.ParseToken(tokenString)
		if err != nil {
			response.New(c).Status(http.StatusUnauthorized).Errors(err).Message("invalid token").MessageID("auth.middleware.failed").Dispatch()
			return
		}
		blackToken, errBlackRepo := blackRepo.FindByToken(tokenString)
		if blackToken != nil && errBlackRepo != nil {
			response.New(c).Status(http.StatusUnauthorized).Errors(err).Message("token is in the black list !").MessageID("auth.middleware.black_list").Dispatch()
			return
		}
		c.Set("claims", token)
		c.Set("user_uid", token.UID)
		ctx := c.Request.Context()
		userUid, err := primitive.ObjectIDFromHex(token.UID)
		ctx = context.WithValue(ctx, "user_uid", userUid)
		c.Request = c.Request.WithContext(ctx)
		c.Set("access_token", tokenString)
		c.Next()
	}
}
