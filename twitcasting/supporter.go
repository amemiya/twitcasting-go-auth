package twitcasting

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SupportingStatusContainer struct {
	IsSupporting bool `json:"is_supporting"`
	Supported    int  `json:"supported"`
	TargetUser   User `json:"target_user"`
}

type SupporterListContainer struct {
	Total      int             `json:"total"`
	Supporting []SupporterUser `json:"supporting"`
}

type SupporterUser struct {
	Id              string `json:"id"`
	ScreenId        string `json:"screen_id"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	Profile         string `json:"profile"`
	Level           int    `json:"level"`
	LastMovieId     string `json:"last_movie_id"`
	IsLive          bool   `json:"is_live"`
	Supported       int    `json:"supported"`
	SupporterCount  int    `json:"supporter_count"`
	SupportingCount int    `json:"supporting_count"`
	Created         int    `json:"created"`
	Point           int    `json:"point"`
	TotalPoint      int    `json:"total_point"`
}

type PostSupportContainer struct {
	AddedCount int `json:"added_count"`
}

type DeleteSupportContainer struct {
	RemovedCount int `json:"removed_count"`
}

type SupporterService ServiceBase

// GetSupportingStatus https://apiv2-doc.twitcasting.tv/#get-supporting-status
func (supporterService *SupporterService) GetSupportingStatus(userId string, targetUserId string, useBearerToken bool) (*SupportingStatusContainer, *ErrorResponse, error) {
	logger := *supporterService.Logger
	response, err := supporterService.Client.get(fmt.Sprintf("/users/%v/supporting_status?target_user_id=%v", userId, targetUserId), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetSupportingStatus", err)
		return nil, nil, err
	}
	defer supporterService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(SupportingStatusContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetSupportingStatus", err)
			return nil, nil, err
		}
		logger.Debug("response for GetSupportingStatus", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetSupportingStatus", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetSupportingStatus", req)
		return nil, req, errors.New("error response")
	}
}

// PostSupport https://apiv2-doc.twitcasting.tv/#support-user
// PostSupport Requests can only be made using a Bearer Token.
func (supporterService *SupporterService) PostSupport(targetUserIds []string) (*PostSupportContainer, *ErrorResponse, error) {
	logger := *supporterService.Logger
	response, err := supporterService.Client.put("/support", map[string]interface{}{"target_user_ids": targetUserIds}, true)
	if err != nil {
		logger.Error("request failed for PostSupport", err)
		return nil, nil, err
	}
	defer supporterService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 || response.StatusCode == 201 {
		req := new(PostSupportContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for PostSupport", err)
			return nil, nil, err
		}
		logger.Debug("response for PostSupport", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for PostSupport", err)
			return nil, nil, err
		}
		logger.Debug("error response for PostSupport", req)
		return nil, req, errors.New("error response")
	}
}

// DeleteSupport https://apiv2-doc.twitcasting.tv/#support-user
// DeleteSupport Requests can only be made using a Bearer Token.
func (supporterService *SupporterService) DeleteSupport(targetUserIds []string) (*DeleteSupportContainer, *ErrorResponse, error) {
	logger := *supporterService.Logger
	response, err := supporterService.Client.put("/unsupport", map[string]interface{}{"target_user_ids": targetUserIds}, true)
	if err != nil {
		logger.Error("request failed for DeleteSupport", err)
		return nil, nil, err
	}
	defer supporterService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(DeleteSupportContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for DeleteSupport", err)
			return nil, nil, err
		}
		logger.Debug("response for DeleteSupport", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for DeleteSupport", err)
			return nil, nil, err
		}
		logger.Debug("error response for DeleteSupport", req)
		return nil, req, errors.New("error response")
	}
}

// GetSupportingList https://apiv2-doc.twitcasting.tv/#supporting-list
func (supporterService *SupporterService) GetSupportingList(userId string, limit int, offset int, useBearerToken bool) (*SupporterListContainer, *ErrorResponse, error) {
	logger := *supporterService.Logger
	response, err := supporterService.Client.get(fmt.Sprintf("/users/%v/supporting?limit=%v&offset=%v", userId, limit, offset), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetSupportingList", err)
		return nil, nil, err
	}
	defer supporterService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(SupporterListContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetSupportingList", err)
			return nil, nil, err
		}
		logger.Debug("response for GetSupportingList", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetSupportingList", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetSupportingList", req)
		return nil, req, errors.New("error response")
	}
}

// GetSupporterList https://apiv2-doc.twitcasting.tv/#supporter-list
func (supporterService *SupporterService) GetSupporterList(userId string, limit int, offset int, sort string, useBearerToken bool) (*SupporterListContainer, *ErrorResponse, error) {
	logger := *supporterService.Logger
	response, err := supporterService.Client.get(fmt.Sprintf("/users/%v/supporters?limit=%v&offset=%v&sort=%v", userId, limit, offset, sort), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetSupporterList", err)
		return nil, nil, err
	}
	defer supporterService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(SupporterListContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetSupporterList", err)
			return nil, nil, err
		}
		logger.Debug("response for GetSupporterList", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetSupporterList", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetSupporterList", req)
		return nil, req, errors.New("error response")
	}
}
