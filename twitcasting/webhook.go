package twitcasting

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type Webhook struct {
	UserId string `json:"user_id"`
	Event  string `json:"event"`
}

type WebhookListContainer struct {
	AllCount int       `json:"all_count"`
	Webhooks []Webhook `json:"webhooks"`
}

type PostWebhookRequestBody struct {
	UserId string   `json:"user_id"`
	Events []string `json:"events"`
}

type PostWebhookContainer struct {
	UserId      string   `json:"user_id"`
	AddedEvents []string `json:"added_events"`
}

type DeleteWebhookContainer struct {
	UserId        string   `json:"user_id"`
	DeletedEvents []string `json:"deleted_events"`
}

type WebhookService ServiceBase

// GetWebhookList https://apiv2-doc.twitcasting.tv/#get-webhook-list
// GetWebhookList Requests can only be made using a Basic Token.
func (webhookService *WebhookService) GetWebhookList(limit int, offset int) (*WebhookListContainer, *ErrorResponse, error) {
	logger := *webhookService.Logger
	response, err := webhookService.Client.get(fmt.Sprintf("/webhooks?limit=%v&offset=%v", limit, offset), false)
	if err != nil {
		logger.Error("request failed for GetWebhookList", err)
		return nil, nil, err
	}
	defer webhookService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(WebhookListContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetWebhookList", err)
			return nil, nil, err
		}
		logger.Debug("response for GetWebhookList", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetWebhookList", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetWebhookList", req)
		return nil, req, errors.New("error response")
	}
}

// PostWebhook https://apiv2-doc.twitcasting.tv/#register-webhook
// PostWebhook Requests can only be made using a Basic Token.
func (webhookService *WebhookService) PostWebhook(userId string, events []string) (*PostWebhookContainer, *ErrorResponse, error) {
	logger := *webhookService.Logger
	response, err := webhookService.Client.post(
		"/webhooks",
		PostWebhookRequestBody{UserId: userId, Events: events},
		false,
	)
	if err != nil {
		logger.Error("request failed for PostWebhook", err)
		return nil, nil, err
	}
	defer webhookService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 || response.StatusCode == 201 {
		req := new(PostWebhookContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for PostWebhook", err)
			return nil, nil, err
		}
		logger.Debug("response for PostWebhook", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for PostWebhook", err)
			return nil, nil, err
		}
		logger.Debug("error response for PostWebhook", req)
		return nil, req, errors.New("error response")
	}
}

// DeleteWebhook https://apiv2-doc.twitcasting.tv/#remove-webhook
// DeleteWebhook Requests can only be made using a Basic Token.
func (webhookService *WebhookService) DeleteWebhook(userId string, events []string) (*DeleteWebhookContainer, *ErrorResponse, error) {
	logger := *webhookService.Logger
	u, err := url.Parse("/webhooks")
	if err != nil {
		logger.Error("parse url failed for DeleteWebhook", err)
		return nil, nil, err
	}
	q := u.Query()
	q.Set("user_id", userId)
	for _, event := range events {
		q.Add("events[]", event)
	}
	u.RawQuery = q.Encode()
	response, err := webhookService.Client.delete(u.String(), false)
	if err != nil {
		logger.Error("request failed for DeleteWebhook", err)
		return nil, nil, err
	}
	defer webhookService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(DeleteWebhookContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for DeleteWebhook", err)
			return nil, nil, err
		}
		logger.Debug("response for DeleteWebhook", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for DeleteWebhook", err)
			return nil, nil, err
		}
		logger.Debug("error response for DeleteWebhook", req)
		return nil, req, errors.New("error response")
	}
}
