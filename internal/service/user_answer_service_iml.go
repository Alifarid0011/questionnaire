package service

import (
	"context"
	"github.com/Alifarid0011/questionnaire-back-end/internal/dto"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type userAnswerServiceImpl struct {
	repo repository.UserAnswerRepository
}

func NewUserAnswerService(repo repository.UserAnswerRepository) UserAnswerService {
	return &userAnswerServiceImpl{repo: repo}
}

// helper: convert DTO -> Model
func toModelAnswers(dtos []dto.AnswerDTO) []models.Answer {
	var result []models.Answer
	for _, d := range dtos {
		result = append(result, models.Answer{
			QuestionID: d.QuestionID,
			Response:   d.Response,
			Score:      0,     // initial score
			IsCorrect:  false, // default
		})
	}
	return result
}

// helper: convert Model -> DTO
func toDTOAnswers(modelsAnswers []models.Answer) []dto.AnswerDTO {
	var result []dto.AnswerDTO
	for _, m := range modelsAnswers {
		result = append(result, dto.AnswerDTO{
			QuestionID: m.QuestionID,
			Response:   m.Response,
		})
	}
	return result
}

func (s *userAnswerServiceImpl) CreateUserAnswer(ctx context.Context, input *dto.UserAnswerDTO) (*dto.UserAnswerDTO, error) {
	answer := &models.UserAnswer{
		ID:        primitive.NewObjectID(),
		QuizID:    input.QuizID,
		UserID:    ctx.Value("user_uid").(primitive.ObjectID),
		Answers:   toModelAnswers(input.Answers),
		Score:     0,
		Appeal:    false,
		CreatedAt: time.Now(),
	}

	if err := s.repo.UserAnswerCreate(ctx, answer); err != nil {
		return nil, err
	}

	return &dto.UserAnswerDTO{
		ID:      answer.ID,
		QuizID:  answer.QuizID,
		UserID:  &answer.UserID,
		Answers: toDTOAnswers(answer.Answers),
	}, nil
}

func (s *userAnswerServiceImpl) GetUserAnswerByID(ctx context.Context, id primitive.ObjectID) (*dto.UserAnswerDTO, error) {
	answer, err := s.repo.UserAnswerFindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserAnswerDTO{
		ID:      answer.ID,
		QuizID:  answer.QuizID,
		UserID:  &answer.UserID,
		Answers: toDTOAnswers(answer.Answers),
	}, nil
}

func (s *userAnswerServiceImpl) GetUserAnswersByQuizID(ctx context.Context, quizID primitive.ObjectID) ([]*dto.UserAnswerDTO, error) {
	answers, err := s.repo.UserAnswerFindByQuizID(ctx, quizID)
	if err != nil {
		return nil, err
	}

	result := make([]*dto.UserAnswerDTO, len(answers))
	for i, ans := range answers {
		result[i] = &dto.UserAnswerDTO{
			ID:      ans.ID,
			QuizID:  ans.QuizID,
			UserID:  &ans.UserID,
			Answers: toDTOAnswers(ans.Answers),
		}
	}
	return result, nil
}

func (s *userAnswerServiceImpl) GetUserAnswersByUserID(ctx context.Context, userID primitive.ObjectID) ([]*dto.UserAnswerDTO, error) {
	answers, err := s.repo.UserAnswerFindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]*dto.UserAnswerDTO, len(answers))
	for i, ans := range answers {
		result[i] = &dto.UserAnswerDTO{
			ID:      ans.ID,
			QuizID:  ans.QuizID,
			UserID:  &ans.UserID,
			Answers: toDTOAnswers(ans.Answers),
		}
	}
	return result, nil
}

func (s *userAnswerServiceImpl) GetUserAnswersByQuizAndUser(ctx context.Context, quizID, userID primitive.ObjectID) ([]*dto.UserAnswerDTO, error) {
	answers, err := s.repo.UserAnswerFindByQuizIDAndUserID(ctx, quizID, userID)
	if err != nil {
		return nil, err
	}

	result := make([]*dto.UserAnswerDTO, len(answers))
	for i, ans := range answers {
		result[i] = &dto.UserAnswerDTO{
			ID:      ans.ID,
			QuizID:  ans.QuizID,
			UserID:  &ans.UserID,
			Answers: toDTOAnswers(ans.Answers),
		}
	}
	return result, nil
}

func (s *userAnswerServiceImpl) SetUserAnswerAppeal(ctx context.Context, uaID primitive.ObjectID, appeal bool) error {
	return s.repo.UserAnswerSetAppeal(ctx, uaID, appeal)
}
