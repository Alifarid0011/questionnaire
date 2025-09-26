package repository

import "github.com/Alifarid0011/questionnaire-back-end/internal/models"

type BlackListTokenRepository interface {
	Store(Token *models.BlackListToken) error
	FindByToken(token string) (*models.BlackListToken, error)
	EnsureIndexes() error
}
