package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Alifarid0011/questionnaire-back-end/config"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"time"
)

type AuthServiceImpl struct {
	userRepo      repository.UserRepository
	casbinRepo    repository.CasbinRepository
	tokenManager  utils.JwtToken
	refreshRepo   repository.RefreshTokenRepository
	blackListRepo repository.BlackListTokenRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
	tokenManager utils.JwtToken,
	refreshRepo repository.RefreshTokenRepository,
	blackListRepo repository.BlackListTokenRepository,
	casbinRepo repository.CasbinRepository,
) AuthService {
	return &AuthServiceImpl{
		userRepo:      userRepo,
		tokenManager:  tokenManager,
		refreshRepo:   refreshRepo,
		blackListRepo: blackListRepo,
		casbinRepo:    casbinRepo,
	}
}

func (s *AuthServiceImpl) Login(req dto.LoginRequest, userAgent *utils.UserAgent) (dto.LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username, context.TODO())
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("user not found: %w", err)
	}
	if errCompareHashAndPassword := utils.CompareHashAndPassword(user.Password, []byte(req.Password)); errCompareHashAndPassword != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}
	accessTokenExpired := time.Now().Add(config.Get.Token.ExpiryAccessToken * time.Minute)
	refreshTokenExpired := time.Now().Add(config.Get.Token.ExpiryRefreshToken * time.Minute)
	accessToken, errGenerateAccessToken := s.tokenManager.GenerateAccessToken(accessTokenExpired.Unix(), user.UID)
	if errGenerateAccessToken != nil {
		return dto.LoginResponse{}, fmt.Errorf("token generation failed: %w", errGenerateAccessToken)
	}
	refreshToken, errGenerateRefreshToken := s.tokenManager.GenerateRefreshToken(refreshTokenExpired.Unix(), user.UID)
	if errGenerateRefreshToken != nil {
		return dto.LoginResponse{}, fmt.Errorf("refresh token generation failed: %w", errGenerateRefreshToken)
	}
	if errRefreshRepo := s.refreshRepo.Store(user.UID, refreshToken, accessToken, 0, userAgent, time.Now(), refreshTokenExpired); errRefreshRepo != nil {
		return dto.LoginResponse{}, fmt.Errorf("storing refresh token failed: %w", errRefreshRepo)
	}
	return dto.LoginResponse{
		AccessToken:         accessToken,
		RefreshToken:        refreshToken,
		UserID:              user.UID,
		AccessTokenExpired:  accessTokenExpired.Unix(),
		RefreshTokenExpired: refreshTokenExpired.Unix(),
	}, nil
}
func (s *AuthServiceImpl) Logout(AccessToken string, userAgent *utils.UserAgent) error {
	token, err := s.refreshRepo.FindByAccessToken(AccessToken)
	if err != nil {
		return err
	}
	blackToken := &models.BlackListToken{
		Token:     token.AccessToken,
		UserAgent: userAgent,
		UserId:    token.UserUid,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(config.Get.Token.ExpiryAccessToken * time.Minute),
	}
	errBlackListRepo := s.blackListRepo.Store(blackToken)
	if errBlackListRepo != nil {
		return errBlackListRepo
	}
	errDeleteByRefreshToken := s.refreshRepo.DeleteByRefreshToken(token.RefreshToken)
	if errDeleteByRefreshToken != nil {
		return errDeleteByRefreshToken
	}
	return nil
}
func (s *AuthServiceImpl) UseRefreshToken(req dto.RefreshRequest, userAgent *utils.UserAgent) (dto.LoginResponse, error) {
	OldRefreshToken, errRefreshRepo := s.refreshRepo.FindByRefreshTokenWithUser(req.RefreshToken)
	if errRefreshRepo != nil {
		return dto.LoginResponse{}, fmt.Errorf("refresh token not found: %w", errRefreshRepo)
	}
	accessTokenExpired := time.Now().Add(config.Get.Token.ExpiryAccessToken * time.Minute)
	refreshTokenExpired := time.Now().Add(config.Get.Token.ExpiryRefreshToken * time.Minute)
	accessToken, errGenerateAccessToken := s.tokenManager.GenerateAccessToken(accessTokenExpired.Unix(), OldRefreshToken.UserUid)
	if errGenerateAccessToken != nil {
		return dto.LoginResponse{}, fmt.Errorf("token generation failed: %w", errGenerateAccessToken)
	}
	NewRefreshTokenString, errGenerateRefreshToken := s.tokenManager.GenerateRefreshToken(refreshTokenExpired.Unix(), OldRefreshToken.User.UID)
	if errGenerateRefreshToken != nil {
		return dto.LoginResponse{}, fmt.Errorf("refresh token generation failed: %w", errGenerateRefreshToken)
	}
	if errRefreshRepoStore := s.refreshRepo.Store(OldRefreshToken.User.UID, NewRefreshTokenString, accessToken, OldRefreshToken.RefreshUseCount+1, userAgent, time.Now(), refreshTokenExpired); errRefreshRepoStore != nil {
		return dto.LoginResponse{}, fmt.Errorf("storing refresh token failed: %w", errRefreshRepoStore)
	} else {
		errDeleteByToken := s.refreshRepo.DeleteByRefreshToken(OldRefreshToken.RefreshToken.RefreshToken)
		if errDeleteByToken != nil {
			return dto.LoginResponse{}, errDeleteByToken
		}
	}
	return dto.LoginResponse{
		AccessToken:         accessToken,
		RefreshToken:        NewRefreshTokenString,
		UserID:              OldRefreshToken.User.UID,
		AccessTokenExpired:  accessTokenExpired.Unix(),
		RefreshTokenExpired: refreshTokenExpired.Unix(),
	}, nil
}

func (s *AuthServiceImpl) hasRole(roles []string, target string) bool {
	for _, r := range roles {
		if r == target {
			return true
		}
	}
	return false
}
