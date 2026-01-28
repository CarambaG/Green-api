package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL = "https://api.green-api.com"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) GetSettings(idInstance, apiToken string) (*SettingsResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/getSettings/%s", BaseURL, idInstance, apiToken)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get settings: %w", err)
	}
	defer resp.Body.Close()

	var result SettingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

func (c *Client) GetStateInstance(idInstance, apiToken string) (*StateResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", BaseURL, idInstance, apiToken)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get state: %w", err)
	}
	defer resp.Body.Close()

	var result StateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

func (c *Client) SendMessage(idInstance, apiToken, chatID, message string) (*MessageResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", BaseURL, idInstance, apiToken)

	payload := SendMessageRequest{
		ChatID:  chatID,
		Message: message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	defer resp.Body.Close()

	var result MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

func (c *Client) SendFileByURL(idInstance, apiToken, chatID, fileURL, fileName, caption string) (*MessageResponse, error) {
	url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", BaseURL, idInstance, apiToken)

	payload := SendFileRequest{
		ChatID:   chatID,
		FileURL:  fileURL,
		FileName: fileName,
		Caption:  caption,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send file: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var result MessageResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
