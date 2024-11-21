package helper

import (
	"errors"
	"fmt"
	"trash_report/repo/record"

	"github.com/go-resty/resty/v2"
)

type GeminiHelper struct {
	client    *resty.Client
	apiURL    string
	apiKey    string
}

func NewGeminiHelper(apiKey string) *GeminiHelper {
	return &GeminiHelper{
		client: resty.New(),
		apiKey: apiKey,
	}
}

func (h *GeminiHelper) AnalyzeReport(title, description, imageURL string) (*record.GeminiResponse, error) {
	resp, err := h.client.R().
		SetQueryParams(map[string]string{
			"title":       title,
			"description": description,
			"image_url":   imageURL,
		}).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", h.apiKey)).
		SetResult(&record.GeminiResponse{}).
		Post(h.apiURL)

	if err != nil {
		return nil, err
	}

	apiResponse := resp.Result().(*record.GeminiResponse)
	if apiResponse.Status != "success" {
		return nil, errors.New("Gemini API failed: " + apiResponse.Status)
	}

	return apiResponse, nil
}
