package repository

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RefreshTokenRepository interface {
	Store(uid primitive.ObjectID, refreshToken string, accessToken string, countOfUsage int, userAgent *utils.UserAgent, creationTime, expiresAt time.Time) error
	FindByRefreshToken(token string) (*models.RefreshToken, error)
	DeleteByUID(uid string) error
	EnsureIndexes() error
	FindByAccessToken(token string) (*models.RefreshToken, error)
	DeleteByRefreshToken(token string) error
	FindByRefreshTokenWithUser(token string) (*models.RefreshTokenWithUser, error)
}
