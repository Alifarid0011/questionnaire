package service

import (
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
)

type AuthService interface {
	Login(req dto.LoginRequest, agent *utils.UserAgent) (dto.LoginResponse, error)
	UseRefreshToken(req dto.RefreshRequest, userAgent *utils.UserAgent) (dto.LoginResponse, error)
	Logout(token string, userAgent *utils.UserAgent) error
}
