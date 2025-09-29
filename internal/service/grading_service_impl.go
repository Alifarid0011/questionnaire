package service

import (
	"context"
	"errors"
	"github.com/Alifarid0011/questionnaire-back-end/internal/models"
	"github.com/Alifarid0011/questionnaire-back-end/internal/repository"
	"github.com/Alifarid0011/questionnaire-back-end/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

type GradingServiceImpl struct {
	quizRepo       repository.QuizRepository
	userAnswerRepo repository.UserAnswerRepository
	apiClient      utils.ShortAnswerAPIClient
}

func NewGradingService(quizRepo repository.QuizRepository, uaRepo repository.UserAnswerRepository, apiClient utils.ShortAnswerAPIClient) GradingService {
	return &GradingServiceImpl{
		quizRepo:       quizRepo,
		userAnswerRepo: uaRepo,
		apiClient:      apiClient,
	}
}

// GradeUserAnswer grades all questions in a user answer concurrently
func (g *GradingServiceImpl) GradeUserAnswer(ctx context.Context, ua *models.UserAnswer) error {
	quiz, err := g.quizRepo.QuizFindByID(ctx, ua.QuizID)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	totalScore := 0.0
	for i, ans := range ua.Answers {
		wg.Add(1)
		go func(i int, ans models.Answer) {
			defer wg.Done()
			q := findQuestionByID(quiz.Questions, ans.QuestionID)
			if q == nil {
				return
			}
			score, isCorrect := 0.0, false
			switch q.Type {
			case "short":
				// call external API
				res, err := g.apiClient.CheckShortAnswer(ctx, ans.Response[0], q.CorrectAnswer, q.KeyWords)
				if err == nil && res.Accepted {
					score = 1.0
					isCorrect = true
				}
			case "radio", "checkbox":
				isCorrect = checkAnswer(ans.Response, q.CorrectAnswer)
				if isCorrect {
					score = 1.0
				}
			}

			mu.Lock()
			ua.Answers[i].Score = score
			ua.Answers[i].IsCorrect = isCorrect
			totalScore += score
			mu.Unlock()
		}(i, ans)
	}
	wg.Wait()

	// save total score
	ua.Score = totalScore
	return g.userAnswerRepo.UserAnswerUpdate(ctx, ua)
}

// ManualGrading allows a grader to override the score of a specific answer
func (g *GradingServiceImpl) ManualGrading(ctx context.Context, uaID primitive.ObjectID, questionID string, newScore float64) error {
	ua, err := g.userAnswerRepo.UserAnswerFindByID(ctx, uaID)
	if err != nil {
		return err
	}

	found := false
	for i, ans := range ua.Answers {
		if ans.QuestionID == questionID {
			ua.Answers[i].Score = newScore
			ua.Answers[i].IsCorrect = newScore > 0
			found = true
			break
		}
	}
	if !found {
		return errors.New("question not found in user answer")
	}

	// recalculate total score
	total := 0.0
	for _, ans := range ua.Answers {
		total += ans.Score
	}
	ua.Score = total

	return g.userAnswerRepo.UserAnswerUpdate(ctx, ua)
}

func (g *GradingServiceImpl) GradeUserAnswerByID(ctx context.Context, uaID primitive.ObjectID) (*models.UserAnswer, error) {
	ua, err := g.userAnswerRepo.UserAnswerFindByID(ctx, uaID)
	if err != nil {
		return nil, err
	}

	if err := g.GradeUserAnswer(ctx, ua); err != nil {
		return nil, err
	}

	return ua, nil
}

func (g *GradingServiceImpl) ManualGradingByID(ctx context.Context, uaID primitive.ObjectID, questionID string, newScore float64) (*models.UserAnswer, error) {
	if err := g.ManualGrading(ctx, uaID, questionID, newScore); err != nil {
		return nil, err
	}

	ua, err := g.userAnswerRepo.UserAnswerFindByID(ctx, uaID)
	if err != nil {
		return nil, err
	}

	return ua, nil
}

func (g *GradingServiceImpl) SetAppeal(ctx context.Context, uaID primitive.ObjectID, appeal bool) error {
	ua, err := g.userAnswerRepo.UserAnswerFindByID(ctx, uaID)
	if err != nil {
		return err
	}

	ua.Appeal = appeal
	return g.userAnswerRepo.UserAnswerUpdate(ctx, ua)
}

// helper
func findQuestionByID(questions []models.Question, id string) *models.Question {
	for _, q := range questions {
		if q.ID == id {
			return &q
		}
	}
	return nil
}

// helper
func checkAnswer(response, correct []string) bool {
	if len(response) != len(correct) {
		return false
	}
	m := make(map[string]struct{})
	for _, c := range correct {
		m[c] = struct{}{}
	}
	for _, r := range response {
		if _, ok := m[r]; !ok {
			return false
		}
	}
	return true
}
