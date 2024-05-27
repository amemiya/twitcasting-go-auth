package twitcasting

import (
	"encoding/json"
	"errors"
	"fmt"
)

type UserContainer struct {
	User            User `json:"User"`
	SupporterCount  int  `json:"supporter_count"`
	SupportingCount int  `json:"supporting_count"`
}

type VerifyCredentialsContainer struct {
	App             App  `json:"App"`
	User            User `json:"User"`
	SupporterCount  int  `json:"supporter_count"`
	SupportingCount int  `json:"supporting_count"`
}

type User struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	ScreenId        string `json:"screen_id"`
	Image           string `json:"image"`
	Profile         string `json:"profile"`
	Level           int    `json:"level"`
	LatestMovieId   string `json:"latest_movie_id"`
	IsLive          bool   `json:"is_live"`
	SupporterCount  int    `json:"supporter_count"`  // @deprecated
	SupportingCount int    `json:"supporting_count"` // @deprecated
	Created         int    `json:"created"`          // @deprecated
}

type App struct {
	ClientId    string `json:"client_id"`
	Name        string `json:"name"`
	OwnerUserId string `json:"owner_user_id"`
}

type UserService ServiceBase

// GetUser @see https://apiv2-doc.twitcasting.tv/#user
func (userService *UserService) GetUser(userId string, useBearerToken bool) (*UserContainer, *ErrorResponse, error) {
	logger := *userService.Logger
	response, err := userService.Client.get(fmt.Sprintf("/users/%v", userId), useBearerToken)
	if err != nil {
		logger.Error("request failed for GetUser", err)
		return nil, nil, err
	}
	defer userService.Client.BodyClose(response.Body) // response.Body.Close()
	if response.StatusCode == 200 {
		req := new(UserContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetUser", err)
			return nil, nil, err
		}
		logger.Debug("response for GetUser", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetUser", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetUser", req)
		return nil, req, errors.New("error response")
	}
}

// GetVerifyCredentials @see https://apiv2-doc.twitcasting.tv/#verify-credentials
// GetVerifyCredentials Requests can only be made using a Bearer Token.
func (userService *UserService) GetVerifyCredentials() (*VerifyCredentialsContainer, *ErrorResponse, error) {
	logger := *userService.Logger
	response, err := userService.Client.get("/verify_credentials", true)
	if err != nil {
		logger.Error("request failed for GetVerifyCredentials", err)
		return nil, nil, err
	}
	defer userService.Client.BodyClose(response.Body)
	if response.StatusCode == 200 {
		req := new(VerifyCredentialsContainer)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode response body failed for GetVerifyCredentials", err)
			return nil, nil, err
		}
		logger.Debug("response for GetVerifyCredentials", req)
		return req, nil, nil
	} else {
		req := new(ErrorResponse)
		err = json.NewDecoder(response.Body).Decode(req)
		if err != nil {
			logger.Error("decode error response body failed for GetVerifyCredentials", err)
			return nil, nil, err
		}
		logger.Debug("error response for GetVerifyCredentials", req)
		return nil, req, errors.New("error response")
	}
}
