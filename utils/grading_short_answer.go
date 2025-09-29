package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Alifarid0011/questionnaire-back-end/config"
	"net/http"
	"time"
)

// ShortAnswerRequest is the request payload for the short answer API
type ShortAnswerRequest struct {
	StudentAnswer   string   `json:"student_answer"`
	AcceptedAnswers []string `json:"accepted_answers"`
	Keywords        []string `json:"keywords,omitempty"`
}

// ShortAnswerResponse is the response from the short answer API
type ShortAnswerResponse struct {
	SimilarityScore float64 `json:"similarity_score"`
	Accepted        bool    `json:"accepted"`
	Details         struct {
		SBERTSimilarity  float64 `json:"sbert_similarity"`
		NLIEntailCount   int     `json:"nli_entail_count"`
		NLIContraCount   int     `json:"nli_contra_count"`
		NLIEntailAvgProb float64 `json:"nli_entail_avg_prob"`
		NLIContraAvgProb float64 `json:"nli_contra_avg_prob"`
		KeywordsPresent  bool    `json:"keywords_present"`
		StudentNegation  bool    `json:"student_negation"`
	} `json:"details"`
}

// ShortAnswerAPIClient handles calls to the short answer grading API
type ShortAnswerAPIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewShortAnswerAPIClient creates a new instance of the client
func NewShortAnswerAPIClient() ShortAnswerAPIClient {
	return ShortAnswerAPIClient{
		BaseURL: config.Get.SemanticAnswerValidator.Url,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// CheckShortAnswer sends the student's short answer to the API and returns the result
func (c *ShortAnswerAPIClient) CheckShortAnswer(ctx context.Context, studentAnswer string, acceptedAnswers, keywords []string) (*ShortAnswerResponse, error) {
	reqBody := ShortAnswerRequest{
		StudentAnswer:   studentAnswer,
		AcceptedAnswers: acceptedAnswers,
		Keywords:        keywords,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/check_answer", c.BaseURL), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var saResp ShortAnswerResponse
	if err := json.NewDecoder(resp.Body).Decode(&saResp); err != nil {
		return nil, err
	}

	return &saResp, nil
}
